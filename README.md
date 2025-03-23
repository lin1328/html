# Go Static Website Generator

A lightweight static website generator built with Go, easily deployable to any static hosting service.

[中文文档](README_zh.md)

## Features

- Easy to customize and extend
- Data provided through JSON API for frontend use
- ICP filing information display support
- Responsive design for mobile and desktop devices
- Clean personal card-style layout
- Customizable icon links and website links

## Deployment Methods

### Quick Start

The project includes a Makefile to simplify building and deployment:

```bash
# Show all available commands
make help

# Build server application
make build

# Build and run server
make run

# Generate static website
make static

# One-step build: all components and generate static site
make all-in-one

# Clean generated files
make clean
```

### Method 1: Local Server

1. Install Go (version 1.16 or higher)
2. Clone this repository
3. Navigate to the project directory
4. Run `make run`
5. Open your browser and visit http://localhost:8080

### Method 2: Static Generation (Recommended for hosting)

The project provides a dedicated tool to generate a fully static HTML website:

1. Run `make static`
2. The static website will be generated in the `dist` directory
3. Upload the contents of the `dist` directory to any static hosting service (GitHub Pages, Netlify, Vercel, etc.)

## Customization

To customize the website, modify the data in `data/pagedata.go`. This file contains all the website's data structures and content.

## Project Structure

- `main.go` - Main server application
- `data/pagedata.go` - Data structures and content
- `cmd/generate/main.go` - Static site generator
- `templates/template.html` - HTML template
- `static/` - Static assets (CSS, images, etc.)
- `dist/` - Generated static website (after running `make static`)

## Copyright Notice

All resources on this site are collected from the internet for learning and communication purposes only, not for commercial use. If there is any infringement, please contact the site administrator and present copyright proof for deletion!

Copyright © 2020-2025 Linqi All Rights Reserved. 