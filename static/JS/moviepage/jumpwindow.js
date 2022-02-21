const btn = document.querySelectorAll('#write')
const jumpwin = document.querySelector('.win')
const x = document.querySelectorAll('#x')
const x1 = document.querySelector('#x1')
const tt = window.localStorage.getItem('imdb');
const token = window.localStorage.getItem('token');
const jumpwin1 = document.querySelector('.win1')
    // console.log(tt);
    // console.log(jumpwin)
    // jumpwin[0].style.display = 'block';
btn[0].addEventListener('click', async() => {
    jumpwin.style.display = 'block'
    alert('请在网页内找到弹窗并输入评论！我还不会做响应式对不起！')
})
btn[1].addEventListener('click', async() => {
    jumpwin1.style.display = 'block'
    alert('请在网页内找到弹窗并输入评论！我还不会做响应式对不起！')
})
x[0].addEventListener('click', async() => {
    jumpwin.style.display = 'none'
})
x[1].addEventListener('click', async() => {
    jumpwin1.style.display = 'none'
})

const send = document.querySelector('#send');
const point = document.querySelector('#point');
const essay = document.querySelector('#essay');
let essayform = new FormData();
let tokenform = new FormData();
tokenform.append('token', token)
essayform.append('movie', tt)
essayform.append('point', '')
essayform.append('essay', '')
send.addEventListener('click', async() => {
    essayform.set('point', point.value)
    essayform.set('essay', essay.value)
    const sendessay = await fetch('http://121.41.120.238:8080/message/essay', {
        method: 'POST',
        body: essayform,
        headers: tokenform
    })
    const essayres = await sendessay.json();
    console.log(essayres.status)
    console.log(essayres.info)
    if (essayres.status === true) {
        alert('成功！')
        location.reload()
    }
    if (essayres.status === false) {
        alert('失败了，你看看你登录了没')
        location.reload()
    }
})


const send1 = document.querySelector('#send1') //第二个发送
const talk = document.querySelector('#talk') //讨论区的输入栏
let talkform = new FormData();
talkform.append('movie', tt)
talkform.append('msg', '')
send1.addEventListener('click', async() => {
    talkform.set('msg', talk.value)
    const sendMsg = await fetch('http://121.41.120.238:8080/message/sendMsg', {
        method: 'POST',
        body: talkform,
        headers: tokenform
    })
    const smres = await sendMsg.json();
    console.log(smres.status)
    console.log(smres.info)
    if (smres.status === true) {
        alert('成功！')
        location.reload()
    }
    if (smres.status === false) {
        alert('失败了，你看看你登录了没')
        location.reload()
    }
})