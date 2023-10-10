# Merancang Database

1. Insert
    - Insert 5 operators pada table operators.
      
      ![Merancang Database](</12_Join-Union-Agregation/screenshots/insert_a.png> "Merancang Database")
   <br>
    - Insert 3 product type.
      
      ![Merancang Database](</12_Join-Union-Agregation/screenshots/insert_b.png> "Merancang Database")
      
   <br>
    - Insert 2 product dengan product type id = 1, dan operators id = 3.
    
    - Insert 3 product dengan product type id = 2, dan operators id = 1.
      
    - Insert 3 product dengan product type id = 3, dan operators id = 4.
      
      ![Merancang Database](</12_Join-Union-Agregation/screenshots/insert_c.png> "Merancang Database")
   <br>
   
    - Insert product description pada setiap product.
      
      ![Merancang Database](</12_Join-Union-Agregation/screenshots/insert_f.png> "Merancang Database")
      
   <br>
    - Insert 5 user pada tabel user.
   
      ![Merancang Database](</12_Join-Union-Agregation/screenshots/insert_g.png> "Merancang Database")
   <br>
   
    - Insert 3 payment methods.
      
      ![Merancang Database](</12_Join-Union-Agregation/screenshots/insert_h.png> "Merancang Database")
   <br>
   
    - Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)
      ![Merancang Database](</12_Join-Union-Agregation/screenshots/insert_i.png> "Merancang Database")
   <br>
   
    - Insert 3 product di masing-masing transaksi.<br>
      ![Merancang Database](</12_Join-Union-Agregation/screenshots/insert_j.png> "Merancang Database")
   <br>

2. Select
   - Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
   ![Merancang Database](</12_Join-Union-Agregation/screenshots/select_a.png> "Merancang Database")
   <br>
   
   - Tampilkan product dengan id = 3.
   ![Merancang Database](</12_Join-Union-Agregation/screenshots/select_b.png> "Merancang Database")
   <br>
   
   - Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
   ![Merancang Database](</12_Join-Union-Agregation/screenshots/select_c.png> "Merancang Database")
   <br>
   
   - Hitung jumlah user / pelanggan dengan status gender Perempuan.
   ![Merancang Database](</12_Join-Union-Agregation/screenshots/select_d.png> "Merancang Database")
   <br>
   
   - Tampilkan data pelanggan dengan urutan sesuai nama abjad
   ![Merancang Database](</12_Join-Union-Agregation/screenshots/select_e.png> "Merancang Database")
   <br>
   
   - Tampilkan 5 data pada data product
   ![Merancang Database](</12_Join-Union-Agregation/screenshots/select_f.png> "Merancang Database")
   <br>

3. Update
   - Ubah data product id 1 dengan nama ‘product dummy’.
   ![Merancang Database](</12_Join-Union-Agregation/screenshots/update_a.png> "Merancang Database")
   <br>
   
   - Update qty = 3 pada transaction detail dengan product id = 1.
   ![Merancang Database](</12_Join-Union-Agregation/screenshots/update_b.png> "Merancang Database")
   <br>

4. Delete
   - Delete data pada tabel product dengan id = 1.
   ![Merancang Database](</12_Join-Union-Agregation/screenshots/delete_a.png> "Merancang Database")
   <br>
   
   - Delete pada pada tabel product dengan product type id 1.
   ![Merancang Database](</12_Join-Union-Agregation/screenshots/delete_b.png> "Merancang Database")
   <br>

# Join, Union, Sub query, Function

1. Gabungkan data transaksi dari user id 1 dan user id 2.
   ![Merancang Database](</12_Join-Union-Agregation/screenshots/join_a.png> "Merancang Database")
   <br>
   
2. Tampilkan jumlah harga transaksi user id 1.
   ![Merancang Database](</12_Join-Union-Agregation/screenshots/join_b.png> "Merancang Database")
   <br>
   
3. Tampilkan total transaksi dengan product type 2.
   ![Merancang Database](</12_Join-Union-Agregation/screenshots/join_c.png> "Merancang Database")
   <br>
   
4. Tampilkan semua field table product dan field name table product type yang saling berhubungan.
   ![Merancang Database](</12_Join-Union-Agregation/screenshots/join_d.png> "Merancang Database")
   <br>
   
5. Tampilkan semua field table transaction, field name table product dan field name table user.
   ![Merancang Database](</12_Join-Union-Agregation/screenshots/join_e.png> "Merancang Database")
   <br>
   
6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.
   ![Merancang Database](</12_Join-Union-Agregation/screenshots/join_6.png> "Merancang Database")
   <br>
   
7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.
   ![Merancang Database](</12_Join-Union-Agregation/screenshots/join_7.png> "Merancang Database")
   <br>
   
8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.
   
   ![Merancang Database](</12_Join-Union-Agregation/screenshots/join_8.png> "Merancang Database")
   <br>
