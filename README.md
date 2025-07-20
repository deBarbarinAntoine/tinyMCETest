# tinyMCETest

## Setup

1. Clone the repository in the directory of your choice:
   ```bash
   git clone https://github.com/deBarbarinAntoine/tinyMCETest
   ```
2. In the root directory, execute the following command:
   ```bash
	go mod tidy && go mod vendor
	```
3. Go to the `ui/assets` directory and execute the following command:
    ```bash
   npm i && npm run build:css
    ```
4. Go back to the root directory and create the `uploads` directory:
	```bash
    mkdir uploads
	```
5. Run the Go server:
	```bash
    go run ./cmd/web
	```
6. Open the browser at those URLs:
   - tinyMCE editor 
       ```
       http://localhost:3030/
       ```
   - DaisyUI test with navbar, sidebar and content card & list
       ```
       http://localhost:3030/menu
       ```