## Alghoritm AI

Jawab : 

Pemilihan Algoritma: Pilih algoritma klasifikasi teks yang sesuai. Beberapa pilihan yang umum digunakan termasuk:

####Naive Bayes:

- Deskripsi: Algoritma Naive Bayes adalah algoritma klasifikasi probabilistik yang didasarkan pada Teorema Bayes. Meskipun sederhana, algoritma ini sering digunakan dalam pemrosesan teks karena kecepatan dan ketepatannya dalam banyak tugas.
- Kelebihan: Cepat dalam pelatihan dan pengujian, baik dalam menangani teks pendek. Efektif dalam mengklasifikasikan dokumen atau kata berdasarkan probabilitas yang dihitung.
- Keterbatasan: Menganggap bahwa fitur dalam teks adalah independen, yang bisa tidak selalu benar dalam kasus nyata. Itu sebabnya disebut "naive" (sederhana).

####Support Vector Machines (SVM):

- Deskripsi: SVM adalah algoritma yang kuat dalam memisahkan data ke dalam dua kelas, yang juga dapat diterapkan pada masalah klasifikasi teks. Ini mencari pemisah (garis atau hiperbidang) terbaik antara dua kelas data.
- Kelebihan: Dapat menangani data yang kompleks dan nonlinier dengan menggunakan kernel SVM. SVM memiliki kemampuan generalisasi yang baik dan dapat mengatasi overfitting.
- Keterbatasan: Pemrosesan teks dengan SVM bisa memakan waktu lebih lama daripada beberapa metode lain. Selain itu, pemilihan kernel yang tepat dapat menjadi tantangan.

####Random Forest:

- Deskripsi: Random Forest adalah algoritma ensemble learning yang menggabungkan sejumlah pohon keputusan untuk membuat prediksi. Dalam pemrosesan teks, ini sering digunakan dengan ekstraksi fitur seperti TF-IDF atau Word Embeddings.
- Kelebihan: Sangat efektif dalam mengatasi masalah overfitting. Mampu mengatasi data dengan banyak fitur (termasuk dalam vektor teks yang besar) dan dapat memberikan perkiraan pentingnya fitur.
- Keterbatasan: Random Forest cenderung lebih kompleks daripada algoritma seperti Naive Bayes, sehingga bisa memerlukan lebih banyak sumber daya komputasi.

####Deep Learning (RNN, CNN, atau Transformer):

- Deskripsi: Model-model deep learning, seperti RNN (Recurrent Neural Network), CNN (Convolutional Neural Network), dan Transformer, sangat canggih dan mampu mengatasi tugas pemrosesan teks yang kompleks. Transformer, khususnya, telah menghadirkan kemajuan besar dalam pemahaman konteks teks.
- Kelebihan: Mampu mengatasi konteks dalam teks dengan baik, memproses teks panjang dengan baik, dan menangani tugas-tugas yang memerlukan pemahaman hierarki dan konteks.
- Keterbatasan: Memerlukan jumlah data yang besar dan sumber daya komputasi yang cukup. Dalam beberapa kasus, mungkin memerlukan waktu lama untuk melatih model. Selain itu, memerlukan pengaturan dan penyetelan parameter yang cermat.