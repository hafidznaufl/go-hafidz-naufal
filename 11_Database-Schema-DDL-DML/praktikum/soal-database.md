# Merancang Database

1. dengan project terserah kalian, praktikkan salah satu workflow serderhana dan buktikan dengan Screenshoot. github flow/gitflow/trunkbase<br>
   **Jawab :**

   ![Merancang Database](</11_Database-Schema-DDL-DML/screenshots/database.png> "Merancang Database")
   <br>

# Data Definition Language

1. Create database alta_online_shop.
   **Jawab :**

   ![Data Definition Language DDL](</11_Database-Schema-DDL-DML/screenshots/1.png> "Data Definition Language DDL")
   <br>

2. Dari schema Olshop yang telah kamu kerjakan di, Implementasikanlah menjadi table pada MySQL.
   **Jawab :**

   - Create table user.
   ![Data Definition Language DDL](</11_Database-Schema-DDL-DML/screenshots/2a.png> "Data Definition Language DDL")
   <br>


   - Create table product, product type, operators, product description, payment_method.
   ![Data Definition Language DDL](</11_Database-Schema-DDL-DML/screenshots/2b.png> "Data Definition Language DDL")
   <br>


   - Create table transaction, transaction detail.
   ![Data Definition Language DDL](</11_Database-Schema-DDL-DML/screenshots/2c.png> "Data Definition Language DDL")
   <br>


3. Create tabel kurir dengan field id, name, created_at, updated_at.
    **Jawab** <br>

    ![Data Definition Language DDL](</11_Database-Schema-DDL-DML/screenshots/3.png> "Data Definition Language DDL")
   <br>


4. Tambahkan ongkos_dasar column di tabel kurir.
    **Jawab** <br>

    ![Data Definition Language DDL](</11_Database-Schema-DDL-DML/screenshots/4.png> "Data Definition Language DDL")
   <br>


5. Tambahkan Rename tabel kurir menjadi shipping.
    **Jawab**
    ![Data Definition Language DDL](</11_Database-Schema-DDL-DML/screenshots/5.png> "Data Definition Language DDL")
   <br>


6. Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.
    **Jawab**
    ![Data Definition Language DDL](</11_Database-Schema-DDL-DML/screenshots/6.png> "Data Definition Language DDL")
   <br>


7. Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.
    **Jawab**

    - 1-to-1: payment method description.
    ![Data Definition Language DDL](</11_Database-Schema-DDL-DML/screenshots/7a.png> "Data Definition Language DDL")
   <br>

    - 1-to-many: user dengan alamat.
    ![Data Definition Language DDL](</11_Database-Schema-DDL-DML/screenshots/7b.png> "Data Definition Language DDL")
   <br>

    - many-to-many: user dengan payment method menjadi user_payment_method_detail.
    ![Data Definition Language DDL](</11_Database-Schema-DDL-DML/screenshots/7c.png> "Data Definition Language DDL")
   <br>