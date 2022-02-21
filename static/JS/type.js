//获得所有电影
let formdata = new FormData();
formdata.append('name', '')
const ms = await fetch("http://121.41.120.238:8080/movie/find", {
    method: 'POST',
    body: formdata
})
const msres = await ms.json();
console.log(msres.information.length);
//获得封面
let picform = new FormData();
picform.append('name', '')
picform.append('heading', 'view')
const pics = await fetch("http://121.41.120.238:8080/movie/find", {
    method: 'POST',
    body: picform
})
const picres = await pics.json();
console.log('!!!')
console.log(picres.information);

//获得评分
picform.set('heading', 'liker')
const points = await fetch("http://121.41.120.238:8080/movie/find", {
    method: 'POST',
    body: picform
})
const pores = await points.json();
console.log('!!!')
console.log(pores.information);


//打印所有电影
const moviebox = document.querySelector('.a4') //获得放电影的div
for (let i = 0; i < msres.information.length; i++) {
    const movie = document.createElement('span'); //大盒子
    movie.id = 'movie'
    const text = document.createElement('div') //放文字的盒子
    const pic = document.createElement('img'); //封面
    pic.src = picres.information[i].picture_1;
    pic.width = '150'
    pic.addEventListener('click', async() => {
        localStorage.setItem('imdb', picres.information[i].imdb)
        window.open('moviepage.html');
    })
    const score = document.createElement('span') //评分
    score.innerHTML = pores.information[i].score
    score.id = 'score';
    const tit = document.createElement('span') //标题
    tit.id = 'title';
    tit.addEventListener('click', async() => {
        localStorage.setItem('imdb', picres.information[i].imdb)
        window.open('moviepage.html');
    })
    tit.innerHTML = msres.information[i].name
    text.appendChild(tit); //装盒
    text.appendChild(score);
    movie.appendChild(pic);
    movie.appendChild(text);
    moviebox.appendChild(movie)
}





//获得选项
const choices = document.getElementById('list').getElementsByTagName('span');
console.log(choices);
choices[0].style.color = 'white'
choices[0].style.backgroundColor = '#51A4D7'
let flag;
for (let i = 0; i < choices.length; i++) { //给每个按钮加指令
    choices[i].addEventListener('click', async() => {
        moviebox.innerHTML = '' //清空上次留下的电影
        for (let a = 0; a < choices.length; a++) {
            choices[a].style.color = 'black'
            choices[a].style.backgroundColor = 'white'
        }
        choices[i].style.backgroundColor = '#51A4D7'
        choices[i].style.color = 'white'
        flag = choices[i].innerHTML


        if (flag == '全部类型') { //把所有电影都拍上来
            for (let i = 0; i < msres.information.length; i++) {
                const movie = document.createElement('span'); //大盒子
                movie.id = 'movie'
                const text = document.createElement('div') //放文字的盒子
                const pic = document.createElement('img'); //封面
                pic.src = picres.information[i].picture_1;
                pic.width = '150'
                pic.addEventListener('click', async() => {
                    localStorage.setItem('imdb', picres.information[i].imdb)
                    window.open('moviepage.html');
                })
                const score = document.createElement('span') //评分
                score.innerHTML = pores.information[i].score
                score.id = 'score';
                const tit = document.createElement('span') //标题
                tit.id = 'title';
                tit.addEventListener('click', async() => {
                    localStorage.setItem('imdb', picres.information[i].imdb)
                    window.open('moviepage.html');
                })
                tit.innerHTML = msres.information[i].name
                text.appendChild(tit); //装盒
                text.appendChild(score);
                movie.appendChild(pic);
                movie.appendChild(text);
                moviebox.appendChild(movie)
            }
        } else {
            let flag1 = 0; //判断是否有该类型的变量
            for (let c = 0; c < msres.information.length; c++) { //遍历每一部电影
                const types = msres.information[c].type.split('')
                for (let b = 0; b < types.length; b += 3) { //遍历该电影的类型
                    console.log(types.slice(b, [b + 2]).join(""))
                    if (types.slice(b, [b + 2]).join("") == flag) { // 判断是否含有该类型
                        flag1 = 1; //有
                        break;
                    }
                    flag1 = 0;
                }
                console.log(msres.information[c]);
                if (flag1 == 1) {
                    const movie = document.createElement('span'); //大盒子
                    movie.id = 'movie'
                    const text = document.createElement('div') //放文字的盒子
                    const pic = document.createElement('img'); //封面
                    pic.src = picres.information[c].picture_1;
                    pic.width = '150'
                    pic.addEventListener('click', async() => {
                        localStorage.setItem('imdb', picres.information[c].imdb)
                        window.open('moviepage.html');
                    })
                    const score = document.createElement('span') //评分
                    score.innerHTML = pores.information[c].score
                    score.id = 'score';
                    const tit = document.createElement('span') //标题
                    tit.id = 'title';
                    tit.addEventListener('click', async() => {
                        localStorage.setItem('imdb', picres.information[c].imdb)
                        window.open('moviepage.html');
                    })
                    tit.innerHTML = msres.information[c].name
                    text.appendChild(tit); //装盒
                    text.appendChild(score);
                    movie.appendChild(pic);
                    movie.appendChild(text);
                    moviebox.appendChild(movie);
                }
                flag1 = 0;
            }
        }
    })
}