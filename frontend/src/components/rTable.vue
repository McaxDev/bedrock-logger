<template>
    <div class="r-table">
        <a-watermark :content="dateTime" :gap="[50,50]" :zIndex="999" :font="{fontSize:12}">
            <a-table class="r-table-atable" :sticky="true" :rowKey="record => record.id" :scroll="{x:true}" :expandedRowKeys="expandedKeys" @expand="toggleExpand" :loading="state.loadingTable" :columns="state.tabelDatas.field" :data-source="state.tabelDatas.data" bordered :pagination="false" :expand-column-width="100" style="max-height: calc(100vh - 40px);flex-direction: column;display: flex;">
                <template #headerCell="{title, column}">
                    <div style="text-wrap: nowrap;">{{ formatTableHeader(title, column) }}</div>
                </template>
                <template #expandColumnTitle>
                    <span>更多</span>
                </template>
                <template #expandedRowRender="{ record, index, indent, expanded }">
                    <div style="margin: 0">
                        <a-table :pagination="false" :data-source="formatMoreInfo(record,'data')" :columns="formatMoreInfo(record,'columns')">
                            <template #bodyCell="{ text, record, index, column }">
                                {{judgeType(text)}}
                            </template>
                        </a-table>
                    </div>
                </template>
                <template #bodyCell="{ text, column }">
                    <span style="text-wrap: nowrap;">{{judgeType(text)}}</span>
                </template>
                <template #title>
                    <div style="display: flex;justify-content: space-between;align-items: center;">
                        <a-typography-text strong>{{ appInfo.main }}</a-typography-text>
                        <div style="display: flex;flex-direction: row;align-items: center;">
                            <a-button type="text" @click="fullScreenTable()">
                                <ExpandOutlined v-if="!isFull"/>
                                <CompressOutlined v-else/>
                            </a-button>
                        </div>
                    </div>
                </template>
                <template #footer>
                    <a-pagination style="text-align: end;" :current="state.tableInfo.current" :total="state.tabelDatas.count" :pageSize="state.tableInfo.pageSize" :showSizeChanger="true" @change="changePage" @showSizeChange="changePageSize"/>
                </template>
            </a-table>
        </a-watermark>
        
    </div>
</template>
  
<script>
import { http } from '@/utils/http'
import { stateStore,useThemeStore,appInfoStore } from '@/stores/stores'
import { ExpandOutlined,CompressOutlined } from '@ant-design/icons-vue'
import dayjs from 'dayjs'
import 'dayjs/locale/zh-cn'
dayjs.locale('zh-cn')

export default{
    components:{ ExpandOutlined,CompressOutlined },
    data(){
        return{
            state:stateStore(),
            isFull:false,
            expandedKeys:[],
            theme:useThemeStore().main,
            dateTime:null,
            appInfo:appInfoStore()
        }
    },
    mounted(){
        this.dateTime = dayjs().format('YYYY-MM-DD HH:mm:ss')
        setInterval(()=>{
            this.dateTime = dayjs().format('YYYY-MM-DD HH:mm:ss')
        },1000)
    },
    methods:{
        formatTableHeader(title, column){
            if(this.state.tabelDatas.field.length>0){
                const data = {...column}
                if(data.dataIndex){
                    return title
                }else if(data.RC_TABLE_INTERNAL_COL_DEFINE){
                    return title[0].children
                }
            }
        },
        toggleExpand(expanded, record) {
            const index = this.expandedKeys.indexOf(record.id)
            if (index === -1) {
                this.expandedKeys.push(record.id)
            } else {
                this.expandedKeys.splice(index, 1)
            }
        },
        changePage(current, size){
            this.state.handleTabelInfo({
                current,
                pageSize:size,
            })
            this.queryInfos()
        },
        changePageSize(current, size){
            this.state.handleTabelInfo({
                current,
                pageSize:size,
            })
            this.queryInfos()
        },
        judgeType(variable){
            const type = Object.prototype.toString.call(variable)
            if(variable===''){
                return ' 没有数据 '
            }else{
                if(type==='[object Boolean]'){
                    if(variable){
                        return '是'
                    }else{
                        return '否'
                    }
                }else if(type==='????'){
                    // return variable
                }
                return variable
            }
        },
        queryInfos(){
            this.state.handleLoadingTable(true)
            const query = {
                ...this.transformObjectTypes(this.state.queryForm),
                page:this.state.tableInfo.current,
                limit:this.state.tableInfo.pageSize,
            }
            if(query.table){
                http.post(`/query/${query.table}`,query)
                    .then(res=>{
                        const field = res.data.data.field.map(item=>{
                            return ({
                                title:item.Title,
                                dataIndex:item.Key
                            })
                        })
                        const data = res.data.data.data
                        const count = res.data.data.count
                        const table = {
                            field,data,count
                        }
                        this.state.handleTabelDatas(table)
                        this.state.handleLoadingTable(false)
                    })
                    .catch(err=>{
                        this.state.handleLoadingTable(false)
                    })
            }
        },
        transformObjectTypes(obj) {
            const arr = ['startX','endX','startY','endY','startZ','endZ']
            for (let key in obj) {
                if(key=='isAsc'){
                    obj[key] = obj[key]=='true'?true:false
                }else if(arr.indexOf(key) != -1){
                    obj[key] = Number(obj[key])
                }
            }
            return obj
        },
        fullScreenTable(){
            this.isFull=!this.isFull
            this.$emit('fullScreenTable', this.isFull)
        },
        formatMoreInfo(obj,type){
            let data={}
            let columns=[]
            for (let key in obj) {
                columns.push({
                    title:key,
                    dataIndex:key,
                })
                data[`${key}`]=obj[key]
            }
            if(type==='columns'){
                return columns
            }else if(type==='data'){
                return [data]
            }
        }
    }
}
</script>

<style scoped>
.r-table{
    height: 100%;
    padding-left: 10px;
}
@media (max-width: 992px) {
    .r-table{
        height: auto;
        padding-left: 0;
    }
}
</style>

<style>
.r-table .ant-table-container{
    max-height: calc(100vh - 160px);
    overflow: auto;
}
</style>