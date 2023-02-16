cookie blog

## TODOs

- [ ] Users and Authentication
    - [ ] `POST /user/login`: Existing user login
    - [ ] `POST /users`: Register a new user
    - [ ] `GET /user`: Get current user
    - [ ] `PUT /user`: Update current user
- [ ] Profiles
    - [ ] `GET /profiles/{username}`: Get a profile
    - [ ] `POST /profiles/{username}/follow`: Follow a user
    - [ ] `DELETE /profiles/{username}/follow`: Unfollow a user
- [ ] Articles
    - [ ] `GET /articles/feed`: Get recent articles from users you follow
    - [x] `GET /articles`: List recent articles globally
    - [x] `POST /articles `: Create an article
    - [x] `GET /articles/:id`: Get an article
    - [x] `PUT /articles/:id`: Update an article
    - [x] `DELETE /articles/:id`: Delete an article
- [ ] Comments
    - [ ] `GET /articles/:id/comments`: Get comments for an article
    - [ ] `POST /articles/:id/comments`: Create a comment for an article
    - [ ] `DELETE /articles/:id/comments/:id`: Delete a comment for an article
- [ ] Clap
    - [ ] `POST /articles/:id/favorite`: Clap an article
    - [ ] `DELETE /articles/:id/favorite`: Unclap an article


- [ ] `GET /tags`: Get tags
- [x] `GET /ping`: Testing server