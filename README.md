users Tablosu İçin Endpointler
GET /users: Tüm kullanıcıları listele.
GET /users/{id}: Belirli bir kullanıcıyı getir.
POST /users: Yeni bir kullanıcı oluştur.
PUT /users/{id}: Belirli bir kullanıcıyı güncelle.
DELETE /users/{id}: Belirli bir kullanıcıyı sil.

posts Tablosu İçin Endpointler
GET /posts: Tüm gönderileri listele.
GET /posts/{id}: Belirli bir gönderiyi getir.
GET /users/{id}/posts: Belirli bir kullanıcının tüm gönderilerini listele.
POST /posts: Yeni bir gönderi oluştur.
PUT /posts/{id}: Belirli bir gönderiyi güncelle.
DELETE /posts/{id}: Belirli bir gönderiyi sil.

likes_dislikes Tablosu İçin Endpointler
GET /likes_dislikes: Tüm beğeni ve beğenmeme kayıtlarını listele.
GET /likes_dislikes/{id}: Belirli bir beğeni veya beğenmeyi getir.
GET /posts/{post_id}/likes_dislikes: Belirli bir gönderinin tüm beğeni ve beğenmeme kayıtlarını listele.
POST /likes_dislikes: Yeni bir beğeni veya beğenmeme oluştur.
PUT /likes_dislikes/{id}: Belirli bir beğeni veya beğenmeyi güncelle.
DELETE /likes_dislikes/{id}: Belirli bir beğeni veya beğenmeyi sil.

commits Tablosu İçin Endpointler
GET /commits: Tüm commit'leri listele.
GET /commits/{id}: Belirli bir commit'i getir.
GET /posts/{post_id}/commits: Belirli bir gönderinin tüm commit'lerini listele.
POST /commits: Yeni bir commit oluştur.
PUT /commits/{id}: Belirli bir commit'i güncelle.
DELETE /commits/{id}: Belirli bir commit'i sil.

categories Tablosu İçin Endpointler
GET /categories: Tüm kategorileri listele.
GET /categories/{id}: Belirli bir kategoriyi getir.
POST /categories: Yeni bir kategori oluştur.
PUT /categories/{id}: Belirli bir kategoriyi güncelle.
DELETE /categories/{id}: Belirli bir kategoriyi sil.



Örnek Endpointler ve HTTP Metodları
Kullanıcılar
GET /users
Tüm kullanıcıları getirir.
GET /users/{id}
Belirli bir kullanıcıyı getirir.
POST /users
Yeni bir kullanıcı oluşturur.
PUT /users/{id}
Belirli bir kullanıcıyı günceller.
DELETE /users/{id}
Belirli bir kullanıcıyı siler.
Gönderiler
GET /posts
Tüm gönderileri getirir.
GET /posts/{id}
Belirli bir gönderiyi getirir.
GET /users/{id}/posts
Belirli bir kullanıcının tüm gönderilerini getirir.
POST /posts
Yeni bir gönderi oluşturur.
PUT /posts/{id}
Belirli bir gönderiyi günceller.
DELETE /posts/{id}
Belirli bir gönderiyi siler.
Beğeni ve Beğenmeme
GET /likes_dislikes
Tüm beğeni ve beğenmeme kayıtlarını getirir.
GET /likes_dislikes/{id}
Belirli bir beğeni veya beğenmeyi getirir.
GET /posts/{post_id}/likes_dislikes
Belirli bir gönderinin tüm beğeni ve beğenmeme kayıtlarını getirir.
POST /likes_dislikes
Yeni bir beğeni veya beğenmeme oluşturur.
PUT /likes_dislikes/{id}
Belirli bir beğeni veya beğenmeyi günceller.
DELETE /likes_dislikes/{id}
Belirli bir beğeni veya beğenmeyi siler.
Commitler
GET /commits
Tüm commit'leri getirir.
GET /commits/{id}
Belirli bir commit'i getirir.
GET /posts/{post_id}/commits
Belirli bir gönderinin tüm commit'lerini getirir.
POST /commits
Yeni bir commit oluşturur.
PUT /commits/{id}
Belirli bir commit'i günceller.
DELETE /commits/{id}
Belirli bir commit'i siler.
Kategoriler
GET /categories
Tüm kategorileri getirir.
GET /categories/{id}
Belirli bir kategoriyi getirir.
POST /categories
Yeni bir kategori oluşturur.
PUT /categories/{id}
Belirli bir kategoriyi günceller.
DELETE /categories/{id}
Belirli bir kategoriyi siler.