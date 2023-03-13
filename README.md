# RPL Blog API

## API Documentation
Link : [Click Here!](https://documenter.getpostman.com/view/25087235/2s93JtPiS7)

## Description
An API for a Blog Application created as the 4th assignment of admin oprec from SE Lab (RPL). This project is created using the Clean Architecture principle where the codes are separated into some parts which are as follows:
- `utils` which is filled by utilities functions and structs used at a lot of parts in the API
- `config` which is filled by functions used to configure stuffs for the API like database connections
- `entity` which is filled with structs used in the API (mainly Database), where in this case is the database entities :
    - User (One to Many : Blog, BlogLike, CommentLike)
    - Blog (One to Many : Comment, BlogLike)
    - Comment (One to Many : CommentLike)
    - BlogLike
    - CommentLike
- `dto` which is filled with structs to accomodate requests from the web according to the needs
- `middleware` which is filled with functions to act as the middle layer from the web requests to the handlers like CORS and Auth
- `controller` which is filled with controller functions that handle the request from web that have passed the middleware and processes it before passing it to the service, also giving the response back after everything is done
-  `service` which is filled with service functions that are connected to the logical business flow of the application which processes the passed request from controller and continue passing it to the repository
- `repository` which is filled with repository functions that are the layer which connects straight to the database where we can modify things in the database using queries

## Tech Stack
- Golang using Gin and GORM (Back-end)
- PostgreSQL (Database)

## Assignment Instruction
Gaylex merupakan calon CEO perusahaan yang dinamai GoBlog (Golang Blog). Ia ingin membuat sebuah aplikasi untuk memposting blog. Sebelum bisa posting, user harus membuat akun dan tentunya login dimana setelah itu ia bisa melihat informasi detail tentang akunnya, termasuk semua postingan blognya. Adapula apabila user tidak sengaja salah mengisi nama, ia bisa mengubahnya. Namun Gaylex yakin pasti ada orang bodoh yang mau mendelete akunnya jadi ia mengijinkannya. Selain itu, semua orang dapat melihat, like, dan comment semua blog yang ada. Apabila seseorang ingin melihat komen, ia bisa membuka detail dari blog tersebut. Skuy bantu Gaylex membuat backendnya!

Project harus menggunakan clean architecture, memiliki dokumentasi, dan di deploy.

## Features
- Supports Authentication for User by Signing In using the registered username/email and password
- Supports Authorization by the registered user roles ("user" or "admin")
- Supports CRUD operations for Users
    - Create (C)
        - Signing Up User
    - Read (R)
        - Get All User (Authorized for Admin)
        - Get a User by Slug (Authorized for User/Admin)
    - Update (U)
        - Edit Self Name (Authorized for User/Admin)
    - Delete
        - Delete Self (Authorized for User/Admin)
- Supports operations for Blogs
    - Create (C)
        - Post a New Blog (Authorized for User/Admin)
    - Read (R)
        - Get All Blogs
        - Get Blog by Slug
- Able to post comments for Blog (Authorized for User/Admin)
- Able to view all comments (Authorized for Admin)
- Able to give like or unlike Blogs and Comments (Authorized for User/Admin)
- Able to see all likes for Blogs or Comments (Authorized for User/Admin)

## Hardships I Felt
- Dividing the handler directory in my 3rd assignment into 3 parts which is the controller, service, and repository at first feels a little foggy to me. But, after succeeding, I finally understood the value and benefit of using Clean Architecture
- At first, I have quite a problem in understanding the authentication and authorization mechanism using JWT which I am quite unfamiliar with as before I've only had experience with Basic, Digest, and Session Auth but not with Token Auth. Now, I understand the merits of using Token Auth rather than using the other methods
- I was confused at the beginning in how to implement the Auth middleware into some routes that needed it. After some searching, I found out that each route can accept the Auth middleware as a parameter and therefore am able to use it well.
- I never tried Auth using Postman and therefore didn't understand at first how to utilize it. But after some trials and errors, I finally discovered the way to utilize Postman for Auth Features

## Things I Learned
- By focusing on this assignment, I've started learning more about Gin and GORM usage in programming using Golang even further outside the ones that I've already learned during the 3rd assignment.
- The better and cleaner Clean Architecture that is implemented in this assignment also make me learn more about the architecture and it's benefit in making my codes more structured and readable.
- Finally, I've also understood the way of using JWT Token Authentication and Authorization in Golang, especially in the Gin framework. Of course including the steps in testing the Authorization feature using Postman.
