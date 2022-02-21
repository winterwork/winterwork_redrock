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
names.length = msres.information.length
scores.length = msres.information.length
console.log(names)

for (let i = 0; i < msres.information.length; i++) {
    for (let a = i + 1; a < msres.information.length; a++) {
        if (scores[i] < scores[a]) {
            let mid = scores[i];
            let midname = names[i];
            names[i] = names[a];
            scores[i] = scores[a];
            names[a] = midname;
            scores[a] = mid;
        }
    }
}
for (let i = 0; i < msres.information.length; i++) {
    let line = document.createElement('hr');
    line.SIZE = '1';
    // console.log(line);
    // cb.className = 'cb';
    line.noshade = "noshade"
    line.color = "#dddddd"
    line.size = '1'
    ranklist.appendChild(line);

    names[i] = msres.information[i].name
    scores[i] = pnres.information[i].score
    const box = document.createElement('div')
    const num = document.createElement('span')
    num.innerHTML = i + 1;
    const title = document.createElement('span')
    title.innerHTML = names[i] = msres.information[i].name
    box.appendChild(num)
    box.appendChild(title)
    ranklist.appendChild(box)


}