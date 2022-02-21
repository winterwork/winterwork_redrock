const btn1 = document.querySelector('#wtl');
const btn2 = document.querySelector('#hld');
const tt = window.localStorage.getItem('imdb');
const token = window.localStorage.getItem('token')
let wtlform = new FormData();
let formdata = new FormData();
formdata.append('token', token)
wtlform.append('movie', tt);
wtlform.append('type', '');
let btn1f, btn2f;

// 判断是否已经想看过
wtlform.set('type', 'W');
const test = await fetch('http://121.41.120.238:8080/message/pointOnly', {
    method: 'POST',
    body: wtlform,
    headers: formdata
})
const testres = await test.json();
console.log(test)
if (testres.info == '好力！') {
    btn1.innerHTML = ' 想 看 ' //修改按钮内部
    const then = await fetch('http://121.41.120.238:8080/message/deletePoint', {
        method: 'POST',
        body: wtlform,
        headers: formdata
    })
    const thenres = await then.json();
    console.log(thenres.info)
    btn1f = 1;
} else if (testres.info == '你已经关注/标注看过这部电影了！') {
    btn1.innerHTML = ' 你已想看 ' //修改按钮内部
    btn1f = 2;
}

//判断是否已看过
wtlform.set('type', 'D');
const test1 = await fetch('http://121.41.120.238:8080/message/pointOnly', {
    method: 'POST',
    body: wtlform,
    headers: formdata
})
const testres1 = await test1.json();
let flag2 = 0;
if (testres1.info == '好力！') {
    btn2.innerHTML = ' 看 过 ' //修改按钮内部
    btn2f = 1;
    flag2 = 1;
    const then1 = await fetch('http://121.41.120.238:8080/message/deletePoint', {
        method: 'POST',
        body: wtlform,
        headers: formdata
    })
    const thenres1 = await then1.json();
    console.log(thenres1.info)
} else if (testres1.info == '你已经关注/标注看过这部电影了！') {
    btn2.innerHTML = '你已看过'
    flag2 = 1;
    btn2f = 2;
}


btn1.addEventListener('click', async() => {
    if (flag2 == '0')
        alert("你是不是没有登录捏")
    wtlform.set('type', 'W');
    if (btn1f == 1) {
        wtlform.set('type', 'W');
        const wanttolook = await fetch('http://121.41.120.238:8080/message/pointOnly', {
            method: 'POST',
            body: wtlform,
            headers: formdata
        })
        const wtlres = await wanttolook.json();
        // alert(wtlres.info);
        btn1f = 2;
        btn1.innerHTML = '你已想看'
    } else if (btn1f == 2) {
        wtlform.set('type', 'W');
        const nowanttolook = await fetch('http://121.41.120.238:8080/message/deletePoint', {
            method: 'POST',
            body: wtlform,
            headers: formdata
        })
        const nowtlres = await nowanttolook.json();
        btn1f = 1;
        btn1.innerHTML = ' 想 看 '

    }

})
btn2.addEventListener('click', async() => {
    if (flag2 == '0')
        alert("你是不是没有登录捏")
    if (btn2f == 1) {
        wtlform.set('type', 'D');
        const wanttolook = await fetch('http://121.41.120.238:8080/message/pointOnly', {
            method: 'POST',
            body: wtlform,
            headers: formdata
        })
        const wtlres = await wanttolook.json();
        // alert(wtlres.info);
        btn2f = 2;
        btn2.innerHTML = '你已看过'
    } else if (btn2f == 2) {
        wtlform.set('type', 'D');
        const nowanttolook = await fetch('http://121.41.120.238:8080/message/deletePoint', {
            method: 'POST',
            body: wtlform,
            headers: formdata
        })
        const nowtlres = await nowanttolook.json();
        // alert(nowtlres.info);
        btn2f = 1;
        btn2.innerHTML = ' 看 过 '

    }


})