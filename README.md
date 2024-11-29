# FORUM-PROJECT
A web platform for collaboration on programming languages and related technologies.

## INTRODUCTION

This project is a forum platform where users can share problems and solutions related to programming languages, algorithms, and software technologies. The project utilizes RESTful APIs for data sharing between users. The backend is developed using Go, while the frontend is created with HTML, CSS, and JavaScript. Docker can be used to easily access the project and run all the services.

The project is divided into three main components:
- **backend**: APIs written in Go
- **frontend**: Interface created with HTML, CSS, and JS
- **deploy**: Deployment operations using Docker

## SETUP

To run the project locally, follow the steps below:

1. **Install Docker**: [Docker Installation Video](https://www.youtube.com/watch?v=iqqDU2crIEQ)
2. **Download all repositories locally**:
   - [Forum Frontend](https://github.com/burakcanhazir/forum-deploy) - Frontend code
   - [Forum Backend](https://github.com/burakcanhazir/burak-forumend) - Backend code
   - [Forum Deployment](https://github.com/burakcanhazir/burakforum) - Deployment and Docker configuration

---

### **Usage Section**
The usage section is quite clear, but a few spelling errors and additional explanations could be added to make parts clearer.

```markdown
## USAGE

1. **Backend Configuration**: 
   Copy the content of the `.env.example` file in the backend repository.

2. **Create .env File**: 
   In the `forum-backend` folder, create a new `.env` file and configure it with the data you copied.

3. **Start with Docker**: 
   Navigate to the `docker-deployment` folder and run the following command in your terminal:
   ```bash
   docker-compose up --build

## License

This project is licensed under the MIT License - see the [License file](LICENSE) for details.

## Contact

If you have any questions or suggestions about the project, feel free to reach out via [hazirburakcan@gmail.com].

## Authors

- [Burakcan Hazir](https://github.com/burakcanhazir)




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
   - [Forum Frontend](https://github.com/burakcanhazir/forum-deploy) - Frontend kodu
   - [Forum Backend](https://github.com/burakcanhazir/burak-forumend) - Backend kodu
   - [Forum Deployment](https://github.com/burakcanhazir/burakforum) - Dağıtım ve Docker yapılandırması

---

### **. Kullanım Bölümü**
Kullanım kısmı çok net, fakat bir kaç yazım hatası ve daha fazla açıklama ekleyerek kısımları daha anlaşılır hale getirebilirsiniz.

```markdown
## KULLANIM

1. **Backend yapılandırması**: 
   Backend reposunda bulunan `.env.example` dosyasının içeriğini kopyalayın.
   
2. **.env Dosyası Oluşturun**:
   `forum-backend` klasöründe yeni bir `.env` dosyası oluşturun ve içeriği kopyaladığınız verilerle yapılandırın.

3. **Docker ile Başlatma**:
   `docker-deployment` klasörüne gidin ve terminalinizde şu komutu çalıştırın:
   ```bash
   docker-compose up --build

## Lisans

Bu proje MIT Lisansı ile lisanslanmıştır - [Lisans dosyasını](LICENSE) inceleyin. 

## İletişim

Proje ile ilgili herhangi bir sorunuz veya öneriniz varsa, [hazirburakcan@gmail.com] adresi ile iletişime geçebilirsiniz.


## Yazarlar

- [Burakcan Hazir](https://github.com/burakcanhazir)


