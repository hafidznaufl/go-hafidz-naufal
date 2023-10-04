# Merancang Database

1. Kode berikut ini dituliskan tanpa mengikuti kebiasaan penulisan yang disarankan.
    - Berapa banyak kekurangan dalam penulisan kode tersebut?
    - Bagian mana saja terjadi kekurangan tersebut?
    - Tuliskan alasan dari setiap kekurangan tersebut. Alasan bisa diberikan dalam bentuk komentar pada kode yang disediakan berikut.
   **Jawab :**

        1.Tidak kurang dari enam kelemahan yang terdapat dalam kode program tersebut.

        2.Ini adalah beberapa kelemahan yang saya temukan dalam kode program tersebut:

        - Penamaan dalam deklarasi struktur variabel user tidak mengikuti konvensi penamaan dalam Go.
        - Nama-nama field dalam struktur user tidak mencerminkan informasi yang benar.
        - Metode getallusers() dan getuserbyid() menggunakan nama yang tidak konsisten dan tidak sesuai dengan konvensi dalam Go.
        - Variabel t dalam struktur userservice kurang deskriptif dan tidak menggambarkan isi variabel dengan jelas.
        - Metode-metode dalam userservice seharusnya menggunakan receiver dengan tipe pointer (*UserService) agar perubahan dapat dilakukan di dalam metode tersebut.
        - Variabel r dalam perulangan tidak memiliki nama yang memiliki makna yang jelas.

        3.Ini adalah alasan untuk setiap kelemahan tersebut:

        - Struktur dan field harus mengikuti konvensi penamaan dalam huruf kapital agar dapat diakses dari luar paket. Nama field juga sebaiknya mencerminkan informasi yang jelas.
        - Nama metode harus dimulai dengan huruf kapital agar dapat diakses dari luar paket. Penggunaan nama yang konsisten membantu dalam memahami fungsionalitas metode.
        - Variabel harus memiliki nama yang menggambarkan isinya dengan jelas agar lebih mudah dipahami.
        - Penggunaan receiver dengan tipe pointer diperlukan agar metode dapat memodifikasi data pada receiver.
        - Memberikan nama yang bermakna pada variabel akan membantu dalam membaca kode dengan lebih baik.

2. 
