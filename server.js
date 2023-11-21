const express=require ("express")
const os= require("os")

const app=express()
app.get("/",(req,res)=>{
        res.send(`hi from ${os.hostname()}`)
})

const port=3001

app.listen(port,()=>console.log(`listening on port ${port}`))