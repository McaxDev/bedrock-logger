<template>
    <div class="l-form">
        <a-spin :indicator="indicator" :spinning="state.loadingTable">
            
            <div class="l-form-group" :style="{ color:theme.main.colorTextSecondary,background:theme.main.colorBgElevated }">
                <div class="l-form-group-title">基础</div>
                <div class="l-form-group-item">
                    <div class="l-form-item-label">查询类型</div>
                    <div style="width: 70%;">
                        <a-select placeholder="请选择查询类型" style="width: 100%;" v-model:value="form.table" :options="option.table" @change="changeTable"></a-select>
                    </div>
                </div>
                <div class="l-form-group-item">
                    <div class="l-form-item-label">起始日期</div>
                    <div style="width: 70%;">
                        <a-range-picker v-model:value="form.time" show-time @change="changeTime"/>
                    </div>
                </div>
                <div class="l-form-group-item">
                    <div class="l-form-item-label">排序</div>
                    <div style="width: 70%;">
                        <a-select placeholder="请选择排序字段" style="width: 65%;" v-model:value="form.order" :options="option.field" @change="changeOrdersField"></a-select>
                        <a-select placeholder="请选择排序类型" style="width: calc(35% - 10px);margin-left: 10px;" v-model:value="form.isAsc" :options="option.order" @change="changeOrdersOder"></a-select>
                    </div>
                </div>
            </div>
            <div v-if="shows.indexOf('x') != -1 || shows.indexOf('y') != -1 ||shows.indexOf('z') != -1" class="l-form-group" :style="{ color:theme.main.colorTextSecondary,background:theme.main.colorBgElevated }">
                <div class="l-form-group-title">坐标区域</div>
                <div v-if="shows.indexOf('x') != -1" class="l-form-group-item">
                    <div class="l-form-item-label">X</div>
                    <div style="width: 80%;">
                        <a-input style="width: calc(50% - 5px);" type="number" v-model:value="form.startX" placeholder="输入起始坐标" />
                        <a-input style="width: calc(50% - 5px);margin-left: 10px" type="number" v-model:value="form.endX" placeholder="输入结束坐标" />
                    </div>
                </div>
                <div v-if="shows.indexOf('y') != -1" class="l-form-group-item">
                    <div class="l-form-item-label">Y</div>
                    <div style="width: 80%;">
                        <a-input style="width: calc(50% - 5px);" type="number" v-model:value="form.startY" placeholder="输入起始坐标" />
                        <a-input style="width: calc(50% - 5px);margin-left: 10px" type="number" v-model:value="form.endY" placeholder="输入结束坐标" />
                    </div>
                </div>
                <div v-if="shows.indexOf('z') != -1" class="l-form-group-item">
                    <div class="l-form-item-label">Z</div>
                    <div style="width: 80%;">
                        <a-input style="width: calc(50% - 5px);" type="number" v-model:value="form.startZ" placeholder="输入起始坐标" />
                        <a-input style="width: calc(50% - 5px);margin-left: 10px" type="number" v-model:value="form.endZ" placeholder="输入结束坐标" />
                    </div>
                </div>
            </div>
            <div v-if="shows.indexOf('block_id') != -1" class="l-form-group" :style="{ color:theme.main.colorTextSecondary,background:theme.main.colorBgElevated }">
                <div class="l-form-group-title">方块</div>
                <div class="l-form-group-item">
                    <div class="l-form-item-label">方块事件</div>
                    <div style="width: 70%;">
                        <a-input type="text" v-model:value="form.blockId" placeholder="输入方块id" />
                    </div>
                </div>
            </div>
            <div v-if="shows.indexOf('player') != -1" class="l-form-group" :style="{ color:theme.main.colorTextSecondary,background:theme.main.colorBgElevated }">
                <div class="l-form-group-title">玩家</div>
                <div class="l-form-group-item">
                    <div class="l-form-item-label">玩家事件</div>
                    <div style="width: 70%;">
                        <a-input type="text" v-model:value="form.player" placeholder="输入玩家名称" />
                    </div>
                </div>
            </div>
            <div v-if="shows.indexOf('dimension') != -1" class="l-form-group" :style="{ color:theme.main.colorTextSecondary,background:theme.main.colorBgElevated }">
                <div class="l-form-group-title">维度</div>
                <div class="l-form-group-item">
                    <div class="l-form-item-label">事件发生维度</div>
                    <div style="width: 70%;">
                        <a-select placeholder="请选择维度" style="width: 100%;" v-model:value="form.dimension" :options="option.dimension" @change="changeDimension"></a-select>
                    </div>
                </div>
            </div>
            <div v-if="shows.indexOf('is_join') != -1" class="l-form-group" :style="{ color:theme.main.colorTextSecondary,background:theme.main.colorBgElevated }">
                <div class="l-form-group-title">服务器</div>
                <div class="l-form-group-item">
                    <div class="l-form-item-label">服务器事件</div>
                    <div style="width: 70%;">
                        <a-select placeholder="请选择状态" style="width: 100%;" v-model:value="form.isJoin" :options="option.is_join" @change="changeJoin"></a-select>
                    </div>
                </div>
            </div>

            <div v-if="this.form.table=='die'" class="l-form-group" :style="{ color:theme.main.colorTextSecondary,background:theme.main.colorBgElevated }">
                <div class="l-form-group-title">死亡事件</div>
                <div class="l-form-group-item">
                    <div class="l-form-item-label">杀手ID</div>
                    <div style="width: 70%;">
                        <a-input type="text" v-model:value="form.killerId" placeholder="输入杀手ID" />
                    </div>
                </div>
                <div class="l-form-group-item">
                    <div class="l-form-item-label">杀手名称</div>
                    <div style="width: 70%;">
                        <a-input type="text" v-model:value="form.killerName" placeholder="输入杀手名称" />
                    </div>
                </div>
                <div class="l-form-group-item">
                    <div class="l-form-item-label">死亡者ID</div>
                    <div style="width: 70%;">
                        <a-input type="text" v-model:value="form.deadId" placeholder="输入亡者ID" />
                    </div>
                </div>
                <div class="l-form-group-item">
                    <div class="l-form-item-label">死亡者名称</div>
                    <div style="width: 70%;">
                        <a-input type="text" v-model:value="form.deadName" placeholder="输入亡者名称" />
                    </div>
                </div>
            </div>

            <a-button style="margin-top: 10px;width: 100%;color: #fff;" type="primary" @click="queryInfos">查询</a-button>

        </a-spin>
    </div>
</template>
  
<script>
import { http } from '@/utils/http'
import { stateStore,useThemeStore } from '@/stores/stores'
export default{
    data(){
        return{
            form:{},
            option:{
                table:[],
                order:[
                    {
                        value: 'true',
                        label: '升序',
                    },
                    {
                        value: 'false',
                        label: '降序',
                    }
                ],
                field:[],
                dimension:[
                    {
                        value: 'minecraft:overworld',
                        label: '主世界',
                    },
                    {
                        value: 'minecraft:nether',
                        label: '地狱',
                    },
                    {
                        value: 'minecraft:the_end',
                        label: '末地',
                    }
                ],
                is_join:[
                    {
                        value: 'true',
                        label: '进服',
                    },
                    {
                        value: 'false',
                        label: '退服',
                    },
                ]
            },
            shows:[],
            state:stateStore(),
        }
    },
    setup() {
        const theme = useThemeStore()
        return {
            theme
        }
    },
    mounted(){
        this.getTables()
    },
    methods:{
        getTables(){
            http.get('/tables')
                .then(res=>{
                    this.option.table = res.data.data.map(item=>{
                        return ({
                            value:item.Name,
                            label:item.Display,
                            field:item.Fields,
                            able:item.Fields.map(item=>item.Key)
                        })
                    })
                })
        },
        queryInfos(){
            this.state.handleLoadingTable(true)
            const query = {
                ...this.form,
                page:this.state.tableInfo.current,
                limit:this.state.tableInfo.pageSize,
            }
            delete query.time
            if(query.table){
                this.state.handleQueryForm({...this.form})
                http.post(`/query/${query.table}`,this.transformObjectTypes(query))
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
        changeOrdersOder(value){
            this.form.isAsc=value
        },
        changeOrdersField(value){
            this.form.order=value
        },
        changeTable(value){
            this.state.handleTabelDatas({count:0,field:[],data:[]})
            this.state.handleTabelInfo({current:1,pageSize:10})
            this.form={}
            this.form.table=value
            const a = {...((this.option.table.filter(item=>item.value==value))[0])}
            this.option.field=a.field.map(item=>{
                return ({
                    value:item.Key,
                    label:item.Title,
                })
            })
            this.shows=a.able
        },
        changeXFilterType(value){
            this.form.XFilterType=value
        },
        changeYFilterType(value){
            this.form.YFilterType=value
        },
        changeZFilterType(value){
            this.form.ZFilterType=value
        },
        changeDimension(value){
            this.form.dimension=value
        },
        changeJoin(value){
            this.form.is_join=value
        },
        changeTime(value){
            this.form.time=value
            const startTimeX = new Date(value[0].$d)
            const endTimeX = new Date(value[1].$d)
            this.form.startTime=startTimeX.toISOString().replace(/T/, ' ').replace(/\..+/, '').substring(0, 19)
            this.form.endTime=endTimeX.toISOString().replace(/T/, ' ').replace(/\..+/, '').substring(0, 19)
        }
    }
}
</script>

<style scoped>
.l-form{
    display: flex;
    flex-direction: column;
    .l-form-item-label{
        font-size: 0.9rem;
    }
    .l-form-item{
        margin-top: 10px;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }
    .l-form-group{
        margin-top: 10px;
        border-radius: 8px;
        padding: 10px;
        width: 100%;
        .l-form-group-item{
            margin-top: 10px;
            width: 100%;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }
        .l-form-group-title{
            font-size: 1rem;
        }
    }
}
</style>