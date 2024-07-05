const express = require('express');
const app = express();
// load .env file
try {
    require('dotenv').config();
} catch (error) {
    console.log(error)
}
const port = process.env.PORT

app.listen(port, function(){
    console.log(`Server is running on port ${port}`);
})
