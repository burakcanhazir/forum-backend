***ENGLİSH***
# FORUM-PROJECT
A website for collaboration on programming languages and similar technologies.

## INTRODUCTION

This project is a forum platform where users share problems and solutions related to programming languages, algorithms, and software technologies. The platform facilitates data sharing between users via RESTful APIs. The backend is written in Go, while the frontend is developed using HTML, CSS, and JavaScript. By using Docker, you can easily run all services.

The project is divided into three main components:
- **Backend**: APIs written in Go
- **Frontend**: User interface developed using HTML, CSS, and JS
- **Deploy**: Deployment processes with Docker

## SETUP

Follow these steps to run the project locally:

1. **Install Docker**: [Docker Installation Video](https://www.youtube.com/watch?v=iqqDU2crIEQ)
2. **Clone all repositories to your local machine**:
   - [Forum Frontend](https://github.com/burakcanhazir/forum-frontend) - Frontend code
   - [Forum Backend](https://github.com/burakcanhazir/forum-backend) - Backend code
   - [Forum Deployment](https://github.com/burakcanhazir/forum-deploy) - Deployment and Docker configuration

```markdown
## USAGE

1. **To run locally**
   - In the terminal, type `git clone https://github.com/burakcanhazir/forum-backend` and `git clone https://github.com/burakcanhazir/forum-frontend`.
   - Go to the `forum-backend` folder, find the `.env.example` file, fill in the necessary fields, and save the file.
   - Rename the `.env.example` file to `.env` and save it.
   - Go to the `forum-backend` and `forum-frontend` directories.
   - Run both files using the command `go run .`.

2. **To run with Docker**
   - You can visit the DOCKERHUB account "bhazir".
   - Clone the deploy folder from GitHub using the command `git clone https://github.com/burakcanhazir/forum-deploy`.
   - Run the following commands in the terminal: `docker pull bhazir/forum-frontend:1.0` and `docker pull bhazir/forum-backend:1.0`.
   - If you have cloned the `forum-backend` repository, find the `.env.example` file and make the necessary changes as indicated. If you haven't cloned it, here is the content of the `.env.example` file. Copy this content into a `.env` file and follow the above commands step by step.

      ```env
      APP_ENV=development
      APP_PORT=8000
      APP_HOST=backend

      DB_DRIVER=sqlite3
      DB_PATH=./your-project.db # You can change this

      JWT_SECRET=your-secret-key   # You can change this
      JWT_EXPIRATION=3600  

      ALLOWED_ORIGINS=http://localhost:8081
      ALLOWED_METHODS=GET,POST,PUT,DELETE
      ALLOWED_HEADERS=Content-Type,Authorization

      RATE_LIMIT=100  
      RATE_LIMIT_WINDOW=60
      ```

   - Rename the `.env.example` file to `.env`.
   - Go to the Docker-compose directory and run the following command in the terminal: `docker-compose up --build`.

**NOTE**
- If an error occurs, go to the `docker-compose.yml` file in the `forum-deploy` folder and check the file paths. Make sure to enter the correct paths from your local machine. Then, return to the `forum-deploy` directory and try running `docker-compose up --build` again.

## License

This project is licensed under the MIT License - see the [LICENSE file](LICENSE) for details.

## Contact

If you have any questions or suggestions regarding the project, feel free to contact me at [hazirburakcan@gmail.com].

## Authors

- [Burakcan Hazir](https://github.com/burakcanhazir)


***TURKİSH***
# FORUM-PROJECT
Programla dilleri ve benzer teknolojiler için yardımlaşma adını yapılmış bir web sitesi

## TANITIM

Bu proje, programlama dilleri, algoritmalar ve yazılım teknolojileri hakkında kullanıcıların problem ve çözümlerini birbirlerine aktardığı bir forum platformudur. 
Projede, RESTful API'ler kullanılarak kullanıcılar arasında veri paylaşımı yapılmaktadır. Backend kısmında Go dili kullanılmıştır, frontend kısmı ise HTML, CSS ve JavaScript ile geliştirilmiştir.
Projeye Docker ile erişim sağlayarak, tüm servisleri kolayca çalıştırabilirsiniz.

Proje üç ana bileşene ayrılmıştır:
- **backend**: Go ile yazılmış API'ler
- **frontend**: HTML, CSS, JS ile yapılmış arayüz
- **deploy**: Docker ile dağıtım işlemleri

## KURULUM

Projeyi lokalinizde çalıştırabilmek için aşağıdaki adımları takip edin:

1. **Docker'ı kurun**: [Docker Kurulum Videosu](https://www.youtube.com/watch?v=iqqDU2crIEQ)
2. **Tüm repoları lokalinizde indirin**:
   - [Forum Frontend](https://github.com/burakcanhazir/forum-frontend) - Frontend kodu
   - [Forum Backend](https://github.com/burakcanhazir/forum-backend) - Backend kodu
   - [Forum Deployment](https://github.com/burakcanhazir/forum-deploy) - Dağıtım ve Docker yapılandırması


```markdown
## KULLANIM

1. **LOCALDE ÇALIŞTIRMAK İÇİN**
      -Terminale sırasıyla git clone https://github.com/burakcanhazir/forum-backend ve git clone https://github.com/burakcanhazir/forum-frontend yazın.
      -Forum-backend klasöründe sizin için hazırlanan .env.example dosyasına gidin ve belirtilen yerleri doldurun ve kaydedin.
      -.env.example dosyasının adını ".env" olarak değiştirin ve kaydedin
      -forum-backend ve forum-frontend dizinlerine gidin.
      -iki dosyayı "go run ." komutu ile çalıştırın.

2. **DOCKER İLE ÇALIŞTIRMAK İÇİN**  
      -    DOCKERHUB  "bhazir" hesabına gidebilirsiniz.
      - git clone https://github.com/burakcanhazir/forum-deploy komutunu kullanarak Githubtan deploy klasörünü çekin
      - docker pull bhazir/forum-frontend:1.0 ve docker pull bhazir/forum-backend:1.0 komutlarını sırasıyla terminale kopyalayıp yapıştırın.
      - eğer forum-backend reposunu indirdiyseniz .env.example dosyasını bulun ve size belirttiğim kısımları değiştirin. Eğer indirmediyseniz sizin için aşağıya .env.example içeriği bırakıyorum. .env isimli dosyaya yapıştırarak yukarıda belirtilen komutları tek tek uygulayın.

            
      ***APP_ENV=development
      APP_PORT=8000
      APP_HOST=backend

      DB_DRIVER=sqlite3
      DB_PATH=./your-project.db #değişiklik yaparbilirsin  -- you can change


      JWT_SECRET=your-secret-key   #değişiklik yaparbilirsin  -- you can change
      JWT_EXPIRATION=3600  
 
      ALLOWED_ORIGINS=http://localhost:8081
      ALLOWED_METHODS=GET,POST,PUT,DELETE
      ALLOWED_HEADERS=Content-Type,Authorization


      RATE_LIMIT=100  
      RATE_LIMIT_WINDOW=60***


      - .env.example dosyasının ismini ".env" olarak değiştirin. 
      - Docker-compose dizinine gidin. Terminale "docker-compose up --build" komutunu kopyalayıp yapışıtırın.

      **NOT**
      - Eğer hata veriyorsa forum-deploy klasöründeki docker compose.yml dosyasına gidin ve dosya yollarına göz atın. Localinizdeki doğru yolları girin. Ve tekrar forum-deploy dizinine gelip tekrar "docker-compose up --build" komutunu kopyalayıp yapışıtırın.


      

## Lisans

Bu proje MIT Lisansı ile lisanslanmıştır - [Lisans dosyasını](LICENSE) inceleyin. 

## İletişim

Proje ile ilgili herhangi bir sorunuz veya öneriniz varsa, [hazirburakcan@gmail.com] adresi ile iletişime geçebilirsiniz.


## Yazarlar

- [Burakcan Hazir](https://github.com/burakcanhazir)


