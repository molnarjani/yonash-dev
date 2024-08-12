# Codebase for my portfolio website

<img width="400" alt="image" src="https://github.com/user-attachments/assets/2ab64f2d-4177-425b-b468-a6085605aa1a">

https://yonash.dev

## Links
- [Project overview](#project-overview)
- [Project structure](#project-structure)
- [Running](#running)
- [Development](#development)

## Project overview
yonash.dev is powered by a simple server-side rendering Go service

Backend:
- Go web server
- [templ](https://github.com/a-h/templ) for template generation

Frontend:
- HTMX + TailwindCSS

## Project structure
Project structure follows [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

```
.
├── cmd/                        -- app entrypoints
│   └── cmd/server/main.go      -- server main entrypoint
└── internal/                   -- internal packaged
    └── web/                    -- static web assets, server side templates 
        ├── assets/
        ├── static/
        └── templates/          -- go templ components
    └── content/                -- content data for pages
```

## Running

> ❗️ Please make sure that you have installed the executable files for all the necessary tools before starting your project. Exactly:
>
> - `Air`: [https://github.com/air-verse/air](https://github.com/air-verse/air)

> - `Templ`: [https://github.com/a-h/templ](https://github.com/a-h/templ)
> - `golangci-lint`: [https://github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint)

To start your project, run the **Gowebly** CLI command in your terminal:

```console
gowebly run
```

## Development

The backend part is located in the `*.go` files in your project folder.

The `./internal/web/templates` folder contains Templ templates that you can use in your frontend part. Also, the `./internal/web/assets` folder contains the `styles.scss` (main styles) and `scripts.js` (main scripts) files.

The `./internal/web/static` folder contains all the static files: icons, images, PWA (Progressive Web App) manifest and other builded/minified assets.
