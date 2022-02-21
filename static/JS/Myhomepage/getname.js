let tokenvalue = window.localStorage.getItem('token'); //获取token的值

let formdata = new FormData();
formdata.append('token', '');

if (tokenvalue === null) { //token存在
} else {
    formdata.set('token', tokenvalue);
}

const a = await fetch("http://121.41.120.238:8080/user/check", {
        method: 'POST',
        headers: formdata
    }) //验证登录状态

const res = await a.json();
// console.log(res.info);
let username = res.info;

if (res.info == '你还没有登录！') {
    alert('你还没有登录！')
}
if (res.info == '你的登录已过期，请重新登录！') {
    alert('你的登录已过期，请重新登录！')
    self.location = 'loginpage.html';
}

document.getElementById('whos').innerHTML = username + '的账号'
document.querySelector('#welcm').innerHTML = '欢迎你！' + username



//get自我介绍

const intro = document.querySelector('.intro');

let idform = new FormData();
idform.append('username', res.info)
const idget = await fetch('http://121.41.120.238:8080/user/getID', {
    method: 'POST',
    body: idform
})
const idgetres = await idget.json();
// console.log(idgetres.info);

let introform = new FormData();
introform.append('id', idgetres.info);

const introtext = await fetch('http://121.41.120.238:8080/homepage/introduce/get', {
    method: 'POST',
    body: introform
})
const introres = await introtext.json();
// console.log('!!')
// console.log(introres)
intro.innerHTML = introres.info



const wantlook = await fetch('http://121.41.120.238:8080/user/getDetail', {
    method: 'POST',
    body: introform
})
const wlres = await wantlook.json();
// console.log(wlres.information);
const wl = wlres.information.WMovie.split('');
const arl = wlres.information.DMovie.split('');
// console.log(wl + '!')


const boxs = document.querySelectorAll('.mb');
//想看
//转单个字符串为n个
let num = 0;
let flag = 0;
let flag1 = wl.length;
for (let i = 0; i < wl.length; i++) {
    if (wl[i] == '/') {
        num++;
        flag = i + 1;
        for (let a = i + 1; a < wl.length; a++) {
            if (wl[a] == '/') {
                flag1 = a;
                break
            }
        }
        let s = wl.slice(flag, [flag1]).join("");
        // console.log(s);
        //查电影封面
        let coverform = new FormData();
        coverform.append('IMDB', s);
        coverform.append('heading', 'view');
        const cover = await fetch('http://121.41.120.238:8080/movie/findByIMDB', {
            body: coverform,
            method: 'POST'
        })
        const coverres = await cover.json();


        //插封面进盒子
        const coverbox = document.createElement('span');
        coverbox.className = 'cbs'
        const pic = document.createElement('img');
        pic.width = '100'
        pic.addEventListener('click', async() => {
            localStorage.setItem('imdb', s)
            window.open('moviepage.html')
        })
        pic.src = coverres.information[0].picture_1;
        coverbox.appendChild(pic);
        boxs[0].appendChild(coverbox);
        flag = flag1;
        flag1 = wl.length;
    }
}
const nums = document.querySelectorAll('.number');
nums[0].innerHTML = '想看' + num + '部'


//看过
//转单个字符串为n个
let num1 = 0;
flag1 = arl.length;
for (let i = 0; i < arl.length; i++) {
    if (arl[i] == '/') {
        num1++;
        flag = i + 1;
        for (let a = i + 1; a < arl.length; a++) {
            if (arl[a] == '/') {
                flag1 = a;
                break
            }
        }
        let s = arl.slice(flag, [flag1]).join("");
        // console.log(s);
        //查电影封面
        let coverform = new FormData();
        coverform.append('IMDB', s);
        coverform.append('heading', 'view');
        const cover = await fetch('http://121.41.120.238:8080/movie/findByIMDB', {
            body: coverform,
            method: 'POST'
        })
        const coverres = await cover.json();


        //插封面进盒子
        const coverbox = document.createElement('span');
        coverbox.className = 'cbs'
        const pic = document.createElement('img');
        pic.width = '100'
        pic.addEventListener('click', async() => {
            localStorage.setItem('imdb', s)
            window.open('moviepage.html')
        })
        pic.src = coverres.information[0].picture_1;
        coverbox.appendChild(pic);
        boxs[1].appendChild(coverbox);
        flag = flag1;
        flag1 = arl.length;
    }
}
nums[1].innerHTML = '看过' + num1 + '部'



//编辑intro
const introchange = document.querySelector('.introchange');
const txt = document.querySelector('#txt');
let first = introres.info
intro.addEventListener('click', async() => {
    intro.style.display = 'none'
    introchange.style.display = 'block'
    if (first == '')
        txt.value = '点击这里编辑你的个人介绍'
    else
        txt.value = first
        // console.log(introchange.value)
})
const btn = introchange.querySelector('#btn');
// console.log(btn)
btn.addEventListener('click', async() => {

    let changeform = new FormData();
    // console.log(txt.value)
    changeform.append('message', txt.value)
    const update = await fetch('http://121.41.120.238:8080/homepage/introduce/update', {
        method: 'POST',
        body: changeform,
        headers: formdata
    })
    const updtres = await update.json();
    // console.log(txt.value)
    intro.style.display = 'block'
    introchange.style.display = 'none'
        // console.log(updtres)
    if (updtres.info == '上传成功') {
        intro.innerHTML = txt.value
        first = txt.value
    }
})


const essays = document.querySelector('.essays')

//获得留言
let essayform = new FormData();
essayform.append('id', idgetres.info)
const essay = await fetch('http://121.41.120.238:8080/message/msg', {
    method: 'POST',
    body: essayform
})
const esres = await essay.json();
console.log(esres.information)

function getLocalTime(nS) {
    return new Date(parseInt(nS) * 1000).toLocaleString().replace(/:\d{1,2}$/, ' ');
} //转时间

for (let i = 0; i < esres.information.length; i++) {
    // let cb = document.createElement('div');
    let line = document.createElement('hr');
    line.SIZE = '1';
    // console.log(line);
    // cb.className = 'cb';
    line.noshade = "noshade"
    line.color = "#dddddd"
    line.size = '1'

    // essays.appendChild(cb)
    essays.appendChild(line)
    console.log(esres.information[i].msg); //改成essay，内容
    console.log(esres.information[i].movie);
    console.log(esres.information[i].type); //类型
    console.log(esres.information[i].point); //评分
    console.log(esres.information[i].time); //时间
    let ttdate = new FormData();
    ttdate.append('IMDB', esres.information[i].movie)
    const search = await fetch('http://121.41.120.238:8080/movie/findByIMDB', {
        method: 'POST',
        body: ttdate
    })
    const seres = await search.json()
    console.log(seres.information[0].name) //电影名
    const box = document.createElement('div'); //big盒子
    box.className = 'box'
    const movie = document.createElement('span'); //电影名
    movie.innerHTML = '《' + seres.information[0].name + '》';
    movie.id = 'movie'
    const time = document.createElement('div') //时间
    time.id = 'time'
    time.innerHTML = getLocalTime(esres.information[i].time);
    const point = document.createElement('span'); //评分
    point.id = 'point'
    const pointval = document.createElement('span');
    point.id = 'point'
    pointval.id = 'pointval'
    point.innerHTML = '评分：';
    pointval.innerHTML = esres.information[i].point;
    point.appendChild(pointval)
    const essay = document.createElement('div');
    essay.id = 'essay'
    essay.innerHTML = esres.information[i].msg;
    box.appendChild(time)
    box.appendChild(movie)
    box.appendChild(point)
    box.appendChild(essay)
    essays.appendChild(box)

}