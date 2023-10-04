# System Design

1. Gambarkan desain ERD dari sistem pembelian tiket kereta api!
      
      ![System Design](</12_System-Design/screenshots/erd.png> "System Design")

2. Gambarkan use case diagram dari sistem pembelian tiket kereta api!

      ![System Design](</12_System-Design/screenshots/use_case.png> "System Design")

3. Terdapat sebuah query pada SQL yaitu SELECT * FROM users; Dengan tujuan yang sama, tuliskan dalam bentuk 

   - Redis

   ```sh
   SCAN 0 MATCH users:*
   ```

   - Neo4j

   ```
   MATCH (u:User)
   RETURN u
   ```

   - Cassandra

   ```
   SELECT * FROM user_data;
   ```