<div id="top"></div>

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/AYDEV-FR/ISEN-API">
    <img src="images/header.jpg" alt="Logo" height="200">
  </a>

<h3 align="center">ISEN REST API</h3>

  <p align="center">
    A simple scrapping API to get all information from Aurion ENT.
    <br />
    <a href="https://api.isen-cyber.ovh"><strong>Play with demo »</strong></a>
    <br />
    <br />
    ·
    <a href="https://github.com/AYDEV-FR/ISEN-API/issues">Report Bug</a>
    ·
    <a href="https://github.com/AYDEV-FR/ISEN-API/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#security">Security concerns</a></li>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

[![Product Name Screen Shot][product-screenshot]](https://example.com)

The goal of this project is to provide a REST API for the Aurion note manager. Although the Aurion software is an available API, it is an option.

This project was born from the Android ISEN application made by @AydevFR during a N3 project (Making an android app). It uses the same operation. The application scrapes the web interface of Aurion and then converts it and renders it on the android application.

## Security

The idea of creating a REST API poses a problem of security and confidentiality because it means that your credentials or your authentication TOKEN are sent to the server hosting the REST API. The server retrieves the pages, parses them and sends your data back to you. The server has therefore knowledge of your data.

2 solutions:

- You make your requests on api.isen-cyber.ovh, but it is a server owned by a student.
- **You can self-host** the API on one of your server.

<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

- [Golang](https://go.dev/)
- [Goquery](https://github.com/PuerkitoBio/goquery)
- [Docker](https://www.docker.com)

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->

## Usage

### Get your token

There are two ways to get your token:

- Use the dedicated API point, in this case **you trust the server** and you send your credentials to it

```
TOKEN=$(curl -s -X POST https://api.isen-cyber.ovh/v1/token --data '{"username":"firstname.lastname","password":"<REDACTED-PASSWORD>"}')
```

- You get your token directly from ISEN's Aurion website like this

```
TOKEN=$(curl -sD - -X POST https://ent-toulon.isen.fr/login --data-raw 'username=firstname.lastname&password=<REDACTED-PASSWORD>' | grep -oP "JSESSIONID=\K([A-Z0-9]*)")
```

### Get your notations

```
curl -X GET https://api.isen-cyber.ovh/v1/notations -H "Token: $TOKEN" | jq
```

### Get your absences

```
curl -X GET https://api.isen-cyber.ovh/v1/absences -H "Token: $TOKEN" | jq
```

### Get your courses

```
curl -X GET https://api.isen-cyber.ovh/v1/planning -H "Token: $TOKEN" | jq
```

_For more examples, please refer to the [Swagger Documentation](openapi.yml)_

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- ROADMAP -->

## Roadmap

- [ ] Add unit test
- [ ] Add planning route
- [x] Add informations route
- [ ] Add teachers scrapping capabilities
- [ ] Add possibility to automaticly calculate notation average

Please feel free to contribute :)

See the [open issues](https://github.com/AYDEV-FR/ISEN-API/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

[@AydevFR](https://twitter.com/AydevFR) - aymeric (dot) deliencourt (at) isen (dot) yncrea (dot) fr

<p align="right">(<a href="#top">back to top</a>)</p>

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/AYDEV-FR/ISEN-API.svg?style=for-the-badge
[contributors-url]: https://github.com/AYDEV-FR/ISEN-API/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/AYDEV-FR/ISEN-API.svg?style=for-the-badge
[forks-url]: https://github.com/AYDEV-FR/ISEN-API/network/members
[stars-shield]: https://img.shields.io/github/stars/AYDEV-FR/ISEN-API.svg?style=for-the-badge
[stars-url]: https://github.com/AYDEV-FR/ISEN-API/stargazers
[issues-shield]: https://img.shields.io/github/issues/AYDEV-FR/ISEN-API.svg?style=for-the-badge
[issues-url]: https://github.com/AYDEV-FR/ISEN-API/issues
[license-shield]: https://img.shields.io/github/license/AYDEV-FR/ISEN-API.svg?style=for-the-badge
[license-url]: https://github.com/AYDEV-FR/ISEN-API/blob/master/LICENSE
[product-screenshot]: images/demo.gif
