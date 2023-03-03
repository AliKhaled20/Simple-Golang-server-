//Global Packages Required for the Project
const path = require("path");
const express = require("express");

//Starting the server
const app = express();

//Serving Static Files 'HTML CSS JS Images'
app.use(express.static(__dirname+"/statics"));

//Serving Index Page
app.get("/",(req,res)=>{
    res.statusCode = 200;
    res.sendFile(path.join(__dirname, '/statics/pages/index.html'));
});

//Serving Game Page
app.get("/game",(req,res)=>{
    res.statusCode = 200;
    res.sendFile(path.join(__dirname,"/statics/pages/game.html"));
});

//Serving endpoint for google
app.get("/gl",(req,res)=>{
    res.statusCode = 307;
    res.redirect("https://google.com");
});

//Serving endpoint for github
app.get("/ghb",(req,res)=>{
    res.statusCode = 307;
    res.redirect("https://github.com");
});

//Serving endpoint for BZU
app.get("/bzu",(req,res)=>{
    res.statusCode = 307;
    res.redirect("https://www.birzeit.edu/ar");
});

//Serving endpoint for error page
app.get("*",(req,res)=>{
    res.statusCode = 404;
    res.sendFile(path.join(__dirname,"/statics/pages/404.html"));
});

//Specifing a port for the server
app.listen(3000 || process.env.PORT,()=>{
    console.log('Listining...');
});