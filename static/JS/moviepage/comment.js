const tt = window.localStorage.getItem('imdb');
const comments = document.querySelector('.ct');
const many = document.getElementById('many');
const mcl = document.querySelector('#mcl');
const mcs = document.querySelector('#mcs');


let commentform = new FormData();
commentform.append('movie', tt)
const comment = await fetch('http://121.41.120.238:8080/message/msgList', {
    method: 'POST',
    body: commentform
})
const cmres = await comment.json();

let long = 0;
for (let i = 0; i < cmres.information.length; i++) {
    // console.log(cmres.information[i]);
    if (cmres.information[i].type == '1') {
        //创建盒子
        long++;
        let cb = document.createElement('div');
        let line = document.createElement('hr');
        line.SIZE = '1';
        // console.log(line);
        cb.className = 'cb';
        line.noshade = "noshade"
        line.color = "#dddddd"
        line.size = '1'
        comments.appendChild(line);
        comments.appendChild(cb);
        const un = document.createElement('span');
        const un1 = document.createElement('span');
        un1.id = 'no';
        const msg = document.createElement('div');
        un1.innerHTML = '第' + long + '楼'
        un.innerHTML = cmres.information[i].username;
        msg.innerHTML = cmres.information[i].msg;
        cb.appendChild(un);
        cb.appendChild(un1);
        cb.appendChild(msg);
        // console.log(cb);

        //在盒子里写评论
        //在盒子里加点赞键

        const tp = document.createElement('span');

        function getLocalTime(nS) {
            return new Date(parseInt(nS) * 1000).toLocaleString().replace(/:\d{1,2}$/, ' ');
        }
        console.log(getLocalTime(cmres.information[i].time))
        tp.innerHTML = getLocalTime(cmres.information[i].time)
        tp.id = 'date'
        cb.appendChild(tp)

        // window.open('search.html')
        // window.open("search.html", "newwindow", "height=100, width=400, toolbar=no, menubar=no, scrollbars=no, resizable=no, location=no, status=no")
    }
}
many.innerHTML = '(全部 ' + long + ' 条)'

let flag = 0;
for (let i = cmres.information.length - 1; i >= 0; i--) {
    // console.log(cmres.information[i]);
    if (cmres.information[i].type == '2') {
        //创建盒子
        flag++;
        long++;
        let cb = document.createElement('div');
        let line = document.createElement('hr');
        line.SIZE = '1';
        // console.log(line);
        cb.id = 'azhe';
        line.noshade = "noshade"
        line.color = "#dddddd"
        line.size = '1'
            // console.log(cb);
        mcs.appendChild(line);
        mcs.appendChild(cb);
        const un = document.createElement('span');
        const un1 = document.createElement('span');
        un1.className = 'no';
        const msg = document.createElement('div');
        un1.innerHTML = '&nbsp&nbsp&nbsp评分：' + cmres.information[i].point;
        un.innerHTML = cmres.information[i].username;
        msg.innerHTML = cmres.information[i].essay;
        cb.appendChild(un);
        cb.appendChild(un1);
        cb.appendChild(msg);
        // console.log(cb);

        //在盒子里写评论
        //在盒子里加点赞键
        const tp = document.createElement('span');
        tp.id = 'date'

        function getLocalTime(nS) {
            return new Date(parseInt(nS) * 1000).toLocaleString().replace(/:\d{1,2}$/, ' ');
        }
        console.log(cmres.information[i])
            // tp.innerHTML = getLocalTime(nS)
        cb.appendChild(tp)
    }
}
// window.open('search.html')
// window.open("search.html", "newwindow", "height=100, width=400, toolbar=no, menubar=no, scrollbars=no, resizable=no, location=no, status=no")

//评分

mcl.innerHTML = '（全部' + flag + '条）'