<template>
    <aside class="side">
        <div class="side-tips">文章分类</div>
        <el-row type="flex" v-for="(item,index) in kindList" :key="index" class="side-kind">
            <el-tag :type=" isActive == item.Kind ? 'danger':'info'" 
                class="side-kind-item"
                @click="btnArticleList(item)">
                {{item.CNName}}(共{{item.Count}}篇)
            </el-tag>
        </el-row>
    </aside>
</template>

<script>
import { getArticleKind } from "@/api/api"
import { Message } from 'element-ui'
export default {
  data () {
    return {
      kindList: [],
      isActive: -1,
    }
  },

  fetch () {
    return getArticleKind().then(res => {
            if (res.data.state === true) {
                this.kindList = res.data.data
            } else {
                Message({
                message: '获取文章类型失败！',
                type: 'error',
                duration: 3 * 1000
                })
            }
        }).catch((err) => console.log(err));
  },

  methods: {
    btnArticleList: function (item) {
      this.isActive = item.Kind
      this.$router.push({name:'list-kind',params:{kind:item.Kind}})
    },
  }
}
</script>