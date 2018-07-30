<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>登录</title>
<script src="https://cdn.bootcss.com/vue/2.4.2/vue.min.js"></script>
<script src="https://cdn.bootcss.com/axios/0.18.0/axios.min.js"></script>
<link rel="stylesheet" href="/static/css/simple-line-icons.css">

<style type="text/css">
body,html{
    margin:0;
    padding:0;
    height:100%;
    background-image:url("/static/img/login_bg.jpg");
    background-size: cover;
}
#cover{
    z-index:-10;
    width:100%;
    height:100%;
    line-height:100%;
    background: inherit;
    -webkit-filter: blur(1px);
    -moz-filter: blur(1px);
    -ms-filter: blur(1px);
    -o-filter: blur(1px);
    filter: blur(20px);
    filter: progid:DXImageTransform.Microsoft.Blur(PixelRadius=20, MakeShadow=false);
}

#container{
    z-index:10;
    position:fixed;
    top:0;
    left:0;
    width:100%;
    height:100%;
    line-height:100%;
}
#login-form{
    z-index:10;
    margin:0 auto;
    width:500px;
    height:320px;
    margin-top:120px;
    background:white;
    background: hsla(0,0%,100%,.25) border-box;
    overflow: hidden;
    border-radius: .3em;
    box-shadow: 0 0 0 1px hsla(0,0%,100%,.3) inset, 0 .5em 1em rgba(0, 0, 0, 0.6);
    text-shadow: 0 1px 1px hsla(0,0%,100%,.3);

}
#login-form::before{
    content: '';
    position: absolute;
    top: 0; right: 0; bottom: 0; left: 0;
    margin: -30px;
    z-index: -1;
    -webkit-filter: blur(20px);
    filter: blur(20px);
}
#login-form input{
    display:block;
    margin:0 auto;
    margin-top:40px;
    width:200px;
    height:20px;
    border:0;
}
#login-form img{
    display:block;
    margin:0 auto;
    width:80px;
    height:auto;
}
#login-form a{
    display:block;
    color:white;
    font-size:30px;
    text-decoration:none;
    width:40px;
    height:40px;
    margin:0 auto;
    margin-top:20px;
}
</style>

</head>
<body>
<div id="cover"></div>

<div id="container">
    <div id="login-form">
        <img src="/static/img/logo.png">
        <input type="input" id="username" placeholder='请输入用户名'/>
        <input type="password" id="pwd" placeholder='请输入密码'/> 
        <a href="#" id="submit"><i class="icon-arrow-right-circle icons"></i></a>
        <a href="#" style="font-size:12px;width:60px">点此注册</a>
    </div>
</div>

<script>
    document.getElementById('submit').onclick = function(e){
        var user = document.getElementById("username").value;
        var pwd  = document.getElementById("pwd").value;
        axios.post('/user/dologin',{"username":user,"password":pwd})
             .then(function (response) {
                console.log(response);
            });
    };
</script>
</body>
</html>
