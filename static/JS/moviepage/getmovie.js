const tt = window.localStorage.getItem('imdb');
const title = document.querySelector('.a3');
const cb = document.querySelector('.cover');
const ms = document.querySelectorAll('.mt');
const bt = document.getElementById('bt');
const btt = document.getElementById('btt');
const bc = document.getElementById('bc');
const movieTitle = document.getElementById('movieTitle');
const mc = document.querySelector('#mc');
const pn = document.querySelector('#pointnum')

// console.log(bt.innerHTML)


//查电影名
let ttform = new FormData();
ttform.append('IMDB', tt);
const tts = await fetch("http://121.41.120.238:8080/movie/findByIMDB", {
    method: 'POST',
    body: ttform
})
const ttres = await tts.json();
title.innerHTML = ttres.information[0].name
bt.innerHTML = ttres.information[0].name + '的剧情简介 · · · · · ·'
mc.innerHTML = ttres.information[0].name + '的影评 · · · · · ·'
bc.innerHTML = ttres.information[0].name + '的讨论区 · · · · · ·'
movieTitle.innerHTML = ttres.information[0].name;


//得到导演和编剧的imdb码
ttform.append('heading', 'member');
const members = await fetch("http://121.41.120.238:8080/movie/findByIMDB", {
    method: 'POST',
    body: ttform
})
const memres = await members.json()
    // console.log(memres.information[0].player); 


//得到导演和编剧
let fmform = new FormData();
fmform.append('IMDB', '');
fmform.set('IMDB', memres.information[0].director + '/' + memres.information[0].scriptwriter);
const findmeb = await fetch('http://121.41.120.238:8080/member/showMember', {
    method: 'POST',
    body: fmform
})
const fmres = await findmeb.json();

//算一共有几个主演
const pls = memres.information[0].player.split('')
    // console.log(pls);
    // console.log(pls.length);
let flag1 = 1;
for (let i = 0; i < pls.length; i++) {
    if (pls[i] == '/') {
        flag1++;
    }
}
// console.log(flag1)



//得到主演
let playerform = new FormData();
playerform.append('IMDB', '');
playerform.set('IMDB', memres.information[0].player)
const playerfs = await fetch('http://121.41.120.238:8080/member/showMember', {
    method: 'POST',
    body: playerform
})
const plyres = await playerfs.json();
// console.log(plyres.information)

//添加演员数量的盒子
const playerBox = document.getElementById('playerBox');
let flag2 = 0;
for (let i = 0; i < flag1; i++) {
    let players = document.createElement('span')
    players.id = 'membs';
    players.addEventListener('click', async() => {
        window.open('memberpage.html');
        localStorage.setItem('member', plyres.information[i].IMDB) //加个token
    })
    players.name = 'ttms';
    players.innerHTML = plyres.information[i].Name;
    playerBox.appendChild(players);
    flag2++;
    if (flag2 < flag1) {
        let fenge = document.createElement('span');
        fenge.style.color = 'rgb(51,119,170)'
        fenge.innerHTML = '/'
        playerBox.appendChild(fenge)
    }
}

const turntomems = document.getElementsByName('ttms')

//改信息
ms[0].innerHTML = fmres.information[0].Name;
ms[1].innerHTML = fmres.information[1].Name;
ms[3].innerHTML = ttres.information[0].type;
ms[4].innerHTML = ttres.information[0].date;
ms[5].innerHTML = ttres.information[0].long + '分钟';
ms[6].innerHTML = ttres.information[0].alias;
ms[7].innerHTML = ttres.information[0].imdb;

ms[0].addEventListener('click', async() => {
    window.open('memberpage.html');
    localStorage.setItem('member', fmres.information[0].IMDB) //加个token
})
ms[1].addEventListener('click', async() => {
        window.open('memberpage.html');
        localStorage.setItem('member', fmres.information[1].IMDB) //加个token
    })
    //查封面
let coverform = new FormData();
coverform.append('heading', 'view');
coverform.append('IMDB', ttres.information[0].imdb);
const cover = await fetch("http://121.41.120.238:8080/movie/findByIMDB", {
    method: 'POST',
    body: coverform
})
const coverres = await cover.json();
// console.log(coverres.information[0].picture_1); //封面
btt.innerHTML = '&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp' + coverres.information[0].brief;

//插入封面
let pic = document.createElement('img');
pic.src = coverres.information[0].picture_1;
pic.width = '150'
cb.appendChild(pic);

// const ttmems = document.getElementsByName('ttms');
// console.log(ttmems)

//剧照
const pt = document.querySelector('#pt');
pt.innerHTML = ttres.information[0].name + '的剧照 · · · · · ·';

const pics = document.querySelector('#pics');
let picform = new FormData();
picform.append('IMDB', tt)
picform.append('heading', 'view')
const pictures = await fetch('http://121.41.120.238:8080/movie/findByIMDB', {
    method: 'POST',
    body: picform
})
const picres = await pictures.json();
// console.log(picres.information[0])

const pic2 = document.createElement('img');
pic2.src = picres.information[0].picture_2;
pic2.width = '300';
pics.appendChild(pic2);

const pic3 = document.createElement('img');
pic3.src = picres.information[0].picture_3;
pic3.width = '300';
pics.appendChild(pic3);

const pic4 = document.createElement('img');
pic4.src = picres.information[0].picture_4;
pic4.width = '300';
pics.appendChild(pic4);

let pointform = new FormData();
pointform.append('IMDB', tt);
pointform.append('heading', 'liker');
const point = await fetch('http://121.41.120.238:8080/movie/findByIMDB', {
    method: 'POST',
    body: pointform
})
const pointres = await point.json();
console.log(pointres.information[0])
pn.innerHTML = pointres.information[0].score

const person = document.getElementById('person')
person.innerHTML = pointres.information[0].comment_num + '人评价'
const pers = document.querySelectorAll('#per')
pers[0].innerHTML = pointres.information[0].num_5 + '人'
pers[1].innerHTML = pointres.information[0].num_4 + '人'
pers[2].innerHTML = pointres.information[0].num_3 + '人'
pers[3].innerHTML = pointres.information[0].num_2 + '人'
pers[4].innerHTML = pointres.information[0].num_1 + '人'