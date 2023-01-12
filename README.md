# 🤏🎨 minatar 🖼️✨
A **min**imal av**atar** generator for turning strings into embeddable base64-encoded images!

## 🤷‍♂️ What's this?
A fun little project that creates an abstract image from a given string. Perfect for when you need a unique image for a user that doesn't have a profile pic yet.

## 🥽 Examples
Click an example to see the corresponding image 

![alice@example.com](https://user-images.githubusercontent.com/96031819/211981802-929c6eca-ca7d-43ce-b17d-9939696c548d.png)

![fred@server.3ck0.com](https://user-images.githubusercontent.com/96031819/211981924-428f9f4b-70d9-4568-91e8-76d770661576.png)

![sing-a-song-of-sixpence](https://user-images.githubusercontent.com/96031819/211981976-88710030-f584-42f0-9355-d74eee6707ef.png)

![192.168.1.1](https://user-images.githubusercontent.com/96031819/211982049-2f0c80be-838a-4001-9055-b3e727b1e3e1.png)

## ⁉️ Why not use gravatar?
Gravatar is great, but sometimes you just want something a little more... unique. Plus, it's a great excuse to play around with Go! 🚀

## ⚙️ How does it work?
It uses a hash of the given string to seed a random number generator, which then creates random circles and squares with random colors. The resulting image is then encoded as a PNG and returned as a base64-encoded string.

## 🛠️ Usage
  Build the project with `go build`
  Run the binary `./minatar`, it starts a local web server listening on port 8080
  Open a web browser and go to `http://localhost:8080/{string}` to get an image

## ✏️ Todo

- [ ] Add support for different shapes (triangles, rectangles, etc.)
- [ ] Add a way to adjust the size of the shapes
- [ ] Add a way to adjust the spacing between the shapes
- [ ] Add tests

## 📜 License

This app is licensed under the [MIT](https://opensource.org/licenses/MIT) software license. 

## 🤝 Contributing

Welcoming contributions to the project! If you have an idea for a new feature or have found a bug, please open an issue on the [GitHub repository](https://github.com/donuts-are-good/minatar). If you want to implement a new feature or fix a bug yourself, please follow these steps:

1.  Fork the repository
2.  Create a new branch for your changes
3.  Make your changes and commit them to your branch
4.  Open a pull request from your branch to the `master` branch of the repository

## 💰 Donate

If you would like to support the development of this project, you can donate to the following addresses:

-   Bitcoin: bc1qg72tguntckez8qy2xy4rqvksfn3qwt2an8df2n
-   Monero: 42eCCGcwz5veoys3Hx4kEDQB2BXBWimo9fk3djZWnQHSSfnyY2uSf5iL9BBJR5EnM7PeHRMFJD5BD6TRYqaTpGp2QnsQNgC

😆👏 Thank you for your support!
