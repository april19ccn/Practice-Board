const { MongoClient } = require("mongodb");

const uri = "mongodb://localhost:27017";

const client = new MongoClient(uri);

const connect = async () => {
    try {
        await client.connect();
        console.log('Connected successfully to MongoDB');
    } catch (err) {
        console.error('Error connecting to MongoDB:', err);
    }
}

const close = async () => {
    try {
        await client.close();
        console.log('Connection to MongoDB closed');
    } catch (err) {
        console.error('Error closing connection to MongoDB:', err);
    }
}

const visitDB = async (methods: any) => {
    try {
        await connect()
        
        return await methods(client)
    } finally {
        // 单一服务，client是在本服务里只创建和生成了一次，之后全局共用，估连接池只有一个
        // 我的理解是多个服务连接同一个MongoDB，才需要控制每个服务的对数据库的连接和关闭
        // await close(); 
    }
}

export default visitDB;