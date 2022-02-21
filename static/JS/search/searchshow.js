const ss = document.getElementById('searchshow')
const search = document.getElementById("search");
const searchBt = document.getElementById('searchBt');
const titles = document.querySelectorAll('.title') //获得标题div

async function last() {
    // alert('?')
    let sv = window.localStorage.getItem('sv');
    ss.innerHTML = '搜索 ' + sv;
    search.value = sv; //保留上个网页的输入值


    //查电影
    let formdata = new FormData();
    formdata.append('name', '')
    formdata.set('name', search.value);
    const mf = await fetch("http://121.41.120.238:8080/movie/find", {
        method: 'POST',
        body: formdata
    })
    const mfres = await mf.json();
    console.log(mfres.information.length);

    let tts = new Array(mfres.information.length)


    const mps = document.querySelectorAll('.mp'); //获得封面div


    for (var flag = 0; flag < mfres.information.length; flag++) {
        tts[flag] = mfres.information[flag].imdb;

        //获得封面
        let coverform = new FormData();
        coverform.append('IMDB', '');
        coverform.append('heading', 'view');
        coverform.set('IMDB', mfres.information[flag].imdb);
        const IMDBs = await fetch("http://121.41.120.238:8080/movie/findByIMDB", {
            method: 'POST',
            body: coverform
        })
        const Isres = await IMDBs.json();
        console.log(Isres.information[0].picture_1); //封面

        //放入封面
        console.log(flag + '!')
        let cover = document.createElement('img');
        cover.src = Isres.information[0].picture_1;
        cover.width = '120';
        cover.className = 'wow'
        mps[flag].appendChild(cover);

        //改变标题
        const titles = document.querySelectorAll('.title')
        console.log(titles);
        titles[flag].innerHTML = mfres.information[flag].name;





    }

    for (let a = 0; a < mfres.information.length; a++) {
        //打开详情页
        titles[a].addEventListener('click', async() => {
            // alert(tts)
            // alert(tts[0])
            // alert(a)
            localStorage.setItem('imdb', tts[a])
            window.open("moviepage.html");
        })
    }


}


last();







searchBt.addEventListener('click', async() => {

        let formdata = new FormData();
        formdata.append('name', '')
        formdata.set('name', search.value);
        localStorage.setItem('sv', search.value); //加个token保存搜索框内的内容


        // for (var value of formdata.values()) {
        //     console.log(value);
        // }


        //查电影
        const mf = await fetch("http://121.41.120.238:8080/movie/find", {
            method: 'POST',
            body: formdata
        })
        const mfres = await mf.json();
        console.log(mfres.information.length);

        // const titles = document.querySelectorAll('.title')//获得标题div
        const mps = document.querySelectorAll('.mp'); //获得封面div

        for (let i = 0; i < 4; i++) {
            //删除上次搜索出的封面和标题
            mps[i].innerHTML = '';
            titles[i].innerHTML = '';




        }

        let tts = new Array(mfres.information.length)


        //循环
        for (var flag = 0; flag < mfres.information.length; flag++) {
            // console.log(mfres.information[flag]);
            console.log(mfres.information[flag].name) //电影标题
            console.log(mfres.information[flag].imdb) //电影tt码
            tts[flag] = mfres.information[flag].imdb;


            // //打开详情页
            // titles[flag].addEventListener('click', async () => {
            //     // alert(tts)
            //     // alert(tts[0])
            //     alert(flag)
            //     localStorage.setItem('imdb', tts[flag])
            //     // window.open("moviepage.html");
            // })


            //获得封面
            let coverform = new FormData();
            coverform.append('IMDB', '');
            coverform.append('heading', 'view');
            coverform.set('IMDB', mfres.information[flag].imdb);
            const IMDBs = await fetch("http://121.41.120.238:8080/movie/findByIMDB", {
                method: 'POST',
                body: coverform
            })
            const Isres = await IMDBs.json();
            console.log(Isres.information[0].picture_1); //封面

            //放入封面
            let cover = document.createElement('img');
            cover.src = Isres.information[0].picture_1;
            cover.width = '120';
            mps[flag].appendChild(cover);

            //改变标题
            console.log(titles);
            titles[flag].innerHTML = mfres.information[flag].name;








        }
        console.log(tts)

        //改变搜索框内容
        let sv = window.localStorage.getItem('sv');
        ss.innerHTML = '搜索 ' + sv;
        localStorage.setItem('sv', search.value);

        alert(mfres.information.length)
        for (let a = 0; a < mfres.information.length; a++) {
            //打开详情页
            titles[a].addEventListener('click', async() => {
                // alert(tts)
                // alert(tts[0])
                // alert(a)
                localStorage.setItem('imdb', tts[a])
                window.open("moviepage.html");
            })
        }


        // t1.innerHTML = mfres.name;
        // t1.addEventListener
        // }
    })
    // for (let i = 0; i < mfres.information.length; i++) {
    //     //存个imdb码的值
    //     titles[i].addEventListener('click', async () => {
    //         localStorage.setItem('imdb', 'mfres.information[i].imdb')
    //     })


// }




// let sv = window.localStorage.getItem('sv');
// // alert(sv);
// ss.innerHTML = '搜索 ' + sv;