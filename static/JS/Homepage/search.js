async function a() {
    const search = document.getElementById("search");
    const searchBt = document.getElementById('searchBt');
    let formdata = new FormData();
    formdata.append('name', '')
}
a();
searchBt.addEventListener('click', async () => {
    localStorage.setItem('sv', search.value);//加个token保存搜索框内的内容
    self.location = 'search.html'//跳转到搜索界面
})