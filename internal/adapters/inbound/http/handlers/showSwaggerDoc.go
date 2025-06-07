package handlers

import (
	"log"
	"net/http"
	"os"
)

type ServeSwaggerFileHandler struct {
}

func NewShowSwaggerDocHandler() *ServeSwaggerFileHandler {
	return &ServeSwaggerFileHandler{}
}

func (h *ServeSwaggerFileHandler) File(w http.ResponseWriter, r *http.Request) {
	openOpenApiFile := "/docs/swagger.yaml"

	absolutePathLookup, err := os.Getwd()

	if err != nil {
		http.Error(w, "Error getting current working directory: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(absolutePathLookup)

	openOpenApiFile = absolutePathLookup + openOpenApiFile

	fileBuffer, err := os.ReadFile(openOpenApiFile)

	if err != nil {
		http.Error(w, "Error reading OpenAPI file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/x-yaml")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(fileBuffer)
}

func (h *ServeSwaggerFileHandler) GetOpenAPISpec(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>
<head>
    <title>API Documentation</title>
    <meta charset="utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">
    <style>
        body { margin: 0; padding: 0; }
    </style>
</head>
<body>
    <div id="redoc-container"></div>
    <script src="https://cdn.jsdelivr.net/npm/redoc@2.1.3/bundles/redoc.standalone.js"></script>
    <script>
        Redoc.init('/api/v1/docs/openapi.yaml', {
            theme: {
                colors: {
                    primary: {
                        main: '#32329f'
                    }
                },
                typography: {
                    fontSize: '14px',
                    lineHeight: '1.5em',
                    code: {
                        fontSize: '13px'
                    },
                    headings: {
                        fontFamily: 'Montserrat, sans-serif',
                        fontWeight: '600'
                    }
                },
                sidebar: {
                    width: '300px'
                }
            },
            scrollYOffset: 60,
            hideDownloadButton: false,
            disableSearch: false,
            hideHostname: false,
            pathInMiddlePanel: true,
            requiredPropsFirst: true,
            sortPropsAlphabetically: true
        }, document.getElementById('redoc-container'));
    </script>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}
