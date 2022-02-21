let RegUn = document.getElementById('RegUn');
let RegPsw = document.getElementById('RegPsw');
let question = document.getElementById('question');
let answer = document.getElementById('answer');

let regbtn = document.getElementById('regbtn');

let formData = new FormData();
formData.append('username', '');//上传用户名input内的值
formData.append('password', '');//上传密码input内的值
formData.append('question', '');//上传密码input内的值
formData.append('answer', '');//上传密码input内的值

regbtn.onclick = () => {
    formData.set('username', RegUn.value);//上传用户名input内的值
    formData.set('password', RegPsw.value);//上传密码input内的值
    formData.set('question', question.value);//上传密码input内的值
    formData.set('answer', answer.value);//上传密码input内的值
    // for (var value of formData.values()) {
    //     console.log(value);
    // }
    if (RegUn.value == '')
        alert('用户名不能为空!')
    else if (RegPsw.value == '')
        alert('密码不能为空!')
    else if (RegPsw.value.length < 8)
        alert('密码长度过短! 密码长度应为8到16位')
    else if (RegPsw.value.length > 16)
        alert('密码长度过长! 密码长度应为8到16位')
    else if (question.value == '')
        alert("请填写密保问题!")
    else if (answer.value == '')
        alert("请填写密保答案!")

    fetch("http://121.41.120.238:8080/user/register", {
        method: 'POST',
        body: formData
    })
        .then(Response => Response.json())
        .then(res => {
            let a = res.info;
            if(a=='注册成功!' ||a=='该账号已经被注册!')
            alert(a);
        })

}


