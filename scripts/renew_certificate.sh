#/bin/bash

source .env
DOMAIN_BASE_URL="https://api.dnsimple.com/v2/$DNSSIMPLE_API_USER_ID/domains/$DNSSIMPLE_DOMAIN_ID"
TMP_FILE_PATH="certrenew"

# retrieve last cert ID
LAST_CERTIFICATE_ID=$(
    curl -q -s -H "Authorization: Bearer $DNSSIMPLE_API_TOKEN" -H "Accept: application/json" \
        $DOMAIN_BASE_URL/certificates | jq '.data[0].id'
)


echo "Retrieving certificates..."
# used for .pem
# download cert and chain
CERT_SERVER=$(
    curl -q -s -H "Authorization: Bearer $DNSSIMPLE_API_TOKEN" -H "Accept: application/json" \
        $DOMAIN_BASE_URL/certificates/$LAST_CERTIFICATE_ID/download | jq -r '.data.server'
)
CERT_CHAIN=$(
    curl -q -s -H "Authorization: Bearer $DNSSIMPLE_API_TOKEN" -H "Accept: application/json" \
        $DOMAIN_BASE_URL/certificates/$LAST_CERTIFICATE_ID/download | jq -r '.data.chain[0]'
)

# used for .key
# download private key
CERT_PRIVATE_KEY=$(
    curl -q -s -H "Authorization: Bearer $DNSSIMPLE_API_TOKEN" -H "Accept: application/json" \
        $DOMAIN_BASE_URL/certificates/$LAST_CERTIFICATE_ID/private_key | jq -r '.data.private_key'
)
echo "Certificates retrieved!"

# write cert files
# certificate
echo "$CERT_SERVER" > "${TMP_FILE_PATH}.pem"
echo >> "${TMP_FILE_PATH}.pem"
echo "$CERT_CHAIN" >> "${TMP_FILE_PATH}.pem"

# private key
echo "$CERT_PRIVATE_KEY" > "${TMP_FILE_PATH}.key"


echo "Installing certificate..."
# install the certs
scp "${TMP_FILE_PATH}.pem" ec2-user@$AWS_EC2_SERVER_IP:/etc/nginx/certs/www_yonash_dev.pem
scp "${TMP_FILE_PATH}.key" ec2-user@$AWS_EC2_SERVER_IP:/etc/nginx/certs/www_yonash_dev.key
echo "Certificate installed!"

echo "Restarting nginx..."
ssh ec2-user@$AWS_EC2_SERVER_IP "sudo systemctl restart nginx"
echo "Done."

# cleanup
rm "${TMP_FILE_PATH}.pem"
rm "${TMP_FILE_PATH}.key"
