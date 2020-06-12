<template>
  <div class="content-article">
    <div v-for="(item,index) in articleList" :key="index">
        <Article :article="item"/>
    </div>
  </div>
</template>

<script>
import { getArticleList } from "@/api/api"
import { Message } from 'element-ui'
import Article from '@/components/article'
export default {
    components:{
      Article
    },

    data () {
        return {
            articleList: [],
        }
    },

    asyncData () {
        return getArticleList({page:1,size:15}).then(res => {
            if (res.data.state === true) {
                return { articleList: res.data.data }
            } else {
                Message({
                    message: res.data.err,
                    type: 'error',
                    duration: 3 * 1000
                })
            }
        }).catch(err => console.log(err));
    },
}
</script>
