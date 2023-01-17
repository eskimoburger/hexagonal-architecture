const mysql = require('mysql2');


const pool = mysql.createPool({
    host: '127.0.0.1',
    user: 'root',
    password: 'w',
    database: 'banking'
  });
  
  // Loop to generate and insert the data
  for (let i = 1; i <= 2000; i++) {
      let name = `customer${i}`;
      let date_of_birth = `1990-01-01`;
      let city = `city${i}`;
      let zipcode = `zip${i}`;
      let status = `active`;
      pool.query(`INSERT INTO customers (name, date_of_birth, city, zipcode, status) VALUES ('${name}', '${date_of_birth}', '${city}', '${zipcode}', '${status}')`, function (error, results, fields) {
          if (error) throw error;
      });
  }