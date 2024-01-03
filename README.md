# CMS

This is a simple CMS built with Go, TS, HTMX, MongoDB, and TypeScript.

## Features

- Create, read, update, and delete content.
- You can create your content model by creating a new struct in `models.go`.
- Dynamic client-side interactions using HTMX.

## Technologies Used

- **Go (Golang)**: The backend of this application is built with Go. Go is known for its simplicity and efficiency, making it a great choice for web development.
- **TypeScript**: TypeScript is used for writing the client-side scripts, enhancing JavaScript with static type definitions.
- **Templ**: Templ is a lightweight, fast, and flexible templating engine for Go, used to generate the HTML pages for this application.
- **HTMX**: HTMX is a small JavaScript library that allows you to access AJAX, CSS Transitions, WebSockets and Server Sent Events directly in HTML, using attributes. It's used in this application to make the user interface more dynamic.
- **MongoDB**: This application uses MongoDB as its database. MongoDB is a source-available cross-platform document-oriented database program. It uses JSON-like documents with optional schemas.
- **Makefile**: This application uses a Makefile for automating build tasks, including the compilation of TypeScript files.
- **entr**: entr is a utility for running arbitrary commands when files change. It's used in this application to automatically regenerate templates, compile TypeScript, and restart the application when files change.

## Prerequisites

- Go
- MongoDB
- Make
- entr
- Node.js and TypeScript (for compiling TypeScript files)

## Installation

1. Clone this repository.
2. Navigate to the project directory.
3. Run `make` to build and run the application.

## Usage

To create a new blog post, navigate to `/createblog` in your web browser and fill out the form.

## License

[MIT](https://choosealicense.com/licenses/mit/)
