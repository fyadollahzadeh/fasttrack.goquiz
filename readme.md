Take-Home task for fasttrack backend engineer interview

## Installation

### Using Dockerfile

1. Clone the repository:

   ```sh
   git clone https://github.com/fyadollahzadeh/fasttrack.goquiz.git
   ```
2. Navigate to the project directory:

   ```sh
   cd fasttrack.goquiz
   ```
3. Ensure there is a `Dockerfile` inside the `config` folder with the necessary instructions to build the application.

4. Build the Docker image:

   ```sh
   docker build -t fasttrack.goquiz -f config/Dockerfile .
   ```
5. Run the Docker container:

   ```sh
   docker run -p 9796:8080 fasttrack.goquiz
   ```

## Notes

1. You can call endpoints by passing "userId" in the header. This can be any value as there is no authentication; it is just a placeholder.
2. Adding quizzes and questions is not supported. There are three different quizzes that are preloaded from the `config/seed.json` file.
3. You can retrieve all the questions of a specific quiz.
4. Submit all the answers at once and see the score in the response.
5. By passing `quizId`, you can see the result and how you rank compared to other users.
6. You can work with the API using Swagger, Postman, etc.
7. There is also a `frontend.html` file which was created by ChatGPT it uses tailwindCSS and some JS function to interact with application , there are some bugs in it. Don't forget to change `API_BASE_URL` when using that.
8. Run the unit tests using:
   ```sh
   go test ./...
   ```

