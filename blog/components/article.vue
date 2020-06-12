<template>
    <div class="article">
        <div class="article-header">
            <div class="article-title">{{article.Title}}</div>
            <div class="article-watch"><nuxt-link 
            class="article-watch-color"
            :to="{name:'detail-id',query:{id:article.ID}}">阅读文章</nuxt-link></div>
        </div>
        <div class="article-body">
            简介：{{article.Introduce}}
        </div>
        <div class="article-footer">
            <el-tag type="info" size="small">作者 : {{article.Author}}</el-tag>
            <el-tag type="danger" size="small">观看次数 : {{article.Watch}}</el-tag>
            <el-tag type="success" size="small">文章类别 : {{article.CNName}}</el-tag>
            <el-tag type="warning" size="small">发布时间 : {{article.AddTime.replace(/T.*/ig, '')}}</el-tag>
        </div>
    </div>
</template>

<script>
export default {
    props:{
        article:{
            type: Object,
            default: function(){
                return {}
            }
        }
    }
}
</script>


