import { Sequelize } from "sequelize";
import { dbConfig } from "../../config/index";

console.log('🚀 ~ :4 ~ dbConfig::==', dbConfig)


const sequelize = new Sequelize({
    dialect: "mysql",
    host: dbConfig.DB_HOST,
    username: dbConfig.DB_USER,
    password: dbConfig.DB_PASSWORD,
    database: dbConfig.DB_NAME,
    logging: true
});




// Verify the database connection and list tables
(async () => {
    try {
        await sequelize.authenticate();
        console.log('Connection has been established successfully.');

        const [results] = await sequelize.query("SHOW TABLES");
        console.log('Tables in the database:', results);
    } catch (error) {
        console.error('Unable to connect to the database:', error);
    }
})();

export default sequelize;

