<a name="readme-top"></a>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <h1 align="center">Golang URL Shortener</h1>

  <p align="center">
    A web server that will look at the path of an incoming web request and determine if it should redirect the user to a new page.
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

The Golang URL Shortener takes data in map or YAML format and creates an HTTP server that will redirect users to the appropriate page. If the data is in a map it will have request path URLs as keys and destination URLs as values. If the data is in YAML it will have a 'path' key that contains the request path URL as a value as well as a 'url' key that will have the destination URL as a value.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- BUILT WITH -->

### Built With

[![Go][go-shield]][go-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

### Prerequisites

-   Golang:
    [https://go.dev/doc/install](https://go.dev/doc/install)

### Installation

1. Clone the repo
    ```sh
    git clone https://github.com/cjsidler/go-url-shortener.git
    ```
2. Build an executable
    ```sh
    go build -o "go-url-shortener.exe" .\main\main.go
    ```
3. Run the executable. Server will be hosted on localhost at port 8080.
    ```sh
    .\go-url-shortener.exe
    ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## License

[![MIT License][license-shield]][license-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

Contact me via email, LinkedIn, or GitHub:

[![Email][gmail-shield]][gmail-url]
[![LinkedIn][linkedin-shield]][linkedin-url]
[![GitHub][github-shield]][github-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

-   [Gophercises](https://gophercises.com/)
-   [README Template](https://github.com/othneildrew/Best-README-Template)
-   [Link Badges](https://shields.io/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->

[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/collinsidler/
[go-shield]: https://img.shields.io/badge/Golang-blue?style=for-the-badge&logo=go&logoColor=white
[go-url]: https://go.dev
[gmail-shield]: https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white
[gmail-url]: mailto:cjsidler@gmail.com
[github-shield]: https://img.shields.io/badge/GitHub-black?style=for-the-badge&logo=github&logoColor=white
[github-url]: https://github.com/cjsidler/go-cli-quizzer
