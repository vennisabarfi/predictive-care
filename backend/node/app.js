const express = require('express');
const app = express();
const event_router = require('./router');

// load .env file
try {
    require('dotenv').config();
} catch (error) {
    console.log(error)
}
const port = process.env.PORT


app.use(event_router);

app.listen(port, function(){
    console.log(`Server is running on port ${port}`);
})
