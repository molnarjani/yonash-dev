# Codebase for yonash.dev

This README file contains all the necessary information about:

- [Project overview](#project-overview)
- [Starting your project](#starting-your-project)
- [Developing your project](#developing-your-project)

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
```

## Starting your project

> ❗️ Please make sure that you have installed the executable files for all the necessary tools before starting your project. Exactly:
>
> - `Air`: [https://github.com/air-verse/air](https://github.com/air-verse/air)

> - `Templ`: [https://github.com/a-h/templ](https://github.com/a-h/templ)
> - `golangci-lint`: [https://github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint)

To start your project, run the **Gowebly** CLI command in your terminal:

```console
gowebly run
```

## Developing your project

The backend part is located in the `*.go` files in your project folder.

The `./internal/web/templates` folder contains Templ templates that you can use in your frontend part. Also, the `./internal/web/assets` folder contains the `styles.scss` (main styles) and `scripts.js` (main scripts) files.

The `./internal/web/static` folder contains all the static files: icons, images, PWA (Progressive Web App) manifest and other builded/minified assets.