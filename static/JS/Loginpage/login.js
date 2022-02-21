let username = document.getElementById('LoginUn');
let password = document.getElementById('LoginPsw');
let loginbtn = document.getElementById('loginbtn');

let formdata = new FormData();
loginbtn.addEventListener('click', loginclick)
formdata.append('username','');
formdata.append('password','');

async function loginclick() {
    formdata.set('username', username.value);
    formdata.set('password', password.value);
    // for (var value of formdata.values()) {
    //     console.log(value);
    // }
    // console.log(formdata)
    const res = await fetch("http://121.41.120.238:8080/user/login", {
        method: 'POST',
        body: formdata
    })//post账户和密码

    const data = await res.json();
    console.log(data.info);
    localStorage.setItem('token', data.info);//添加token
    
    //实现登录验证

    let checktoken = new FormData()
    checktoken.append('token','');
    checktoken.set('token',data.info);

    const token = data.info
    ;
    if(data.info == '账号或密码错误！')
    alert('账号或密码错误！')
    
    const aaa = await fetch("http://121.41.120.238:8080/user/check", {
        method: 'POST',
        headers: checktoken
    })
    const atext = await aaa.json();
    console.log(atext.info);

    if(atext !='账号或密码错误！')
    {
        alert('登录成功')
        self.location='homepage.html'
    }
}

