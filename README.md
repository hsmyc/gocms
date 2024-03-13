# CMS

This is a simple CMS built with Go and TypeScript.

## Technologies Used

- **Go (Golang)**: The backend of this application is built with Go. Go is known for its simplicity and efficiency, making it a great choice for web development.
- **TypeScript**: TypeScript is used for writing the client-side scripts, enhancing JavaScript with static type definitions.
- **Templ**: Templ is a lightweight, fast, and flexible templating engine for Go, used to generate the HTML pages for this application.
- **MongoDB**: This application uses MongoDB as its database. MongoDB is a source-available cross-platform document-oriented database program. It uses JSON-like documents with optional schemas.
- **Makefile**: This application uses a Makefile for automating build tasks, including the compilation of TypeScript files.
- **entr**: entr is a utility for running arbitrary commands when files change. It's used in this application to automatically regenerate templates, compile TypeScript, and restart the application when files change.
- **Esbuild**: Esbuild is a JavaScript bundler and minifier. It's used in this application to compile TypeScript files.

## Prerequisites

- Go
- MongoDB
- Esbuild
- Make
- entr
- Node.js and TypeScript (for compiling TypeScript files)

## Installation

1. Clone this repository.
2. Navigate to the project directory.
3. Create .env file with the for the following environment variables:
   - `MONGO_URI`: The URI for your MongoDB database.
4. Run 'go mod tidy' to install the required Go packages.
5. Make sure GOPATH and GOROOT are set properly.

   - `GOROOT` is the location where Go is installed.
   - `GOPATH` is the location where Go projects are stored.

   ```bash
     export GOPATH=/Users/yourusername/go
     export PATH=$PATH:$GOPATH
     export PATH=$PATH:$GOPATH/bin
   ```

6. Make sure MongoDB is running.
7. Run `make` to build and run the application.

## License

[MIT](https://choosealicense.com/licenses/mit/)
