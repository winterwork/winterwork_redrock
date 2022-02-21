const ranklist = document.querySelector('.ranklist')
let formdata = new FormData();
formdata.append('name', '')
const ms = await fetch("http://121.41.120.238:8080/movie/find", {
    method: 'POST',
    body: formdata
})
const msres = await ms.json();

formdata.append('heading', 'liker')
const point = await fetch("http://121.41.120.238:8080/movie/find", {
    method: 'POST',
    body: formdata
})
const pnres = await point.json();
console.log(msres.information.length);
let names = new Array();
let scores = new Array();
let imdbs = new Array();
names.length = msres.information.length
scores.length = msres.information.length
imdbs.length = msres.information.length
for (let i = 0; i < msres.information.length; i++) {
    names[i] = msres.information[i].name
    scores[i] = pnres.information[i].score
    imdbs[i] = pnres.information[i].imdb
}
// console.log(names)
// console.log(scores)

for (let i = 0; i < msres.information.length; i++) {
    for (let a = i + 1; a < msres.information.length; a++) {
        if (scores[i] < scores[a]) {
            let mid = scores[i];
            let midname = names[i];
            let midimdb = imdbs[i];
            names[i] = names[a];
            imdbs[i] = imdbs[a];
            scores[i] = scores[a];
            names[a] = midname;
            imdbs[a] = midimdb;
            scores[a] = mid;
        }
    }
}
console.log(scores)
console.log(names)

for (let i = 0; i < msres.information.length; i++) {
    let line = document.createElement('hr');
    line.SIZE = '1';
    line.noshade = "noshade"
    line.color = "#dddddd"
    line.size = '1'
    ranklist.appendChild(line);


    const box = document.createElement('div')
    const num = document.createElement('span')
    num.innerHTML = i + 1;
    const title = document.createElement('span')
    title.innerHTML = names[i]
    title.id = 'title'
    title.addEventListener('click', async() => {
        localStorage.setItem('imdb', imdbs[i])
        window.open("moviepage.html");
    })
    box.appendChild(num)
    box.appendChild(title)
    ranklist.appendChild(box)


}

//推荐
const pics = document.getElementsByTagName('img')
pics[2].addEventListener('click', async() => {
    localStorage.setItem('imdb', 'tt11219254')
    window.open("moviepage.html");
})

pics[3].addEventListener('click', async() => {
    localStorage.setItem('imdb', 'tt11286314')
    window.open("moviepage.html");
})