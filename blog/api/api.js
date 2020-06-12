import axios from 'axios'

// let prodURL = "http://111.229.249.89:8888/blog",
//     devURL = "http://localhost:8080/blog";

let prodURL = "http://111.229.249.89:8888/blog",
    devURL = "http://111.229.249.89:8888/blog";


// http://111.229.249.89:8888/blog/getArticleList
// http://111.229.249.89:8888/blog/getArticleListByKind
// http://111.229.249.89:8888/blog/getArticleByID
// http://111.229.249.89:8888/blog/getArticleKind
// http://111.229.249.89:8888/blog/getBloggerInfo


const request = axios.create({
  baseURL: process.env.NODE_ENV === 'development'? devURL : prodURL,
  timeout: 5000
})


export function getArticleKind() {
    return request({
        url: `/getArticleKind`,
        method: 'get'
    })
}

export function getArticleList(data = {page:1,size:15}) {
    return request({
        url: `/getArticleList?page=${data.page}&size=${data.size}`,
        method: 'get'
    })
}

export function getArticleByID(data = {id:1}) {
    return request({
        url: `/getArticleByID?id=${data.id}`,
        method: 'get'
    })
}

export function getArticleListByKind(data = {page:1,size:15,kind:1}) {
    return request({
        url: `/getArticleListByKind?page=${data.page}&size=${data.size}&kind=${data.kind}`,
        method: 'get'
    })
}

export function getBloggerInfo(data = {id:1}) {
    return request({
        url: `/getBloggerInfo?id=${data.id}`,
        method: 'get'
    })
}
