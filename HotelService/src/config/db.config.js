const dotenv = require('dotenv');

dotenv.config();

const config = {
  development: {
    username: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_NAME,
    host: process.env.DB_HOST,
    port: process.env.DB_PORT,
    dialect: 'mysql',
  
  }
}

console.log('🚀 ~ :22 ~ config::==', config)





module.exports = config;