<template>
	<nocloud-table
		:loading="loading"
		:items="tableData"
		:value="selected"
		@input="handleSelect"
		:single-select="singleSelect"
		:footer-error="fetchError"
	>

	</nocloud-table>
</template>

<script>
import noCloudTable from "@/components/table.vue"

export default {
	name: "accounts-table",
	components: {
		"nocloud-table": noCloudTable
	},
	props: {
		value: {
			type: Array,
			default: () => []
		},
		"single-select": {
			type: Boolean,
			default: false
		}
	},
	data () {
		return {
			selected: this.value,
			loading: false,
			fetchError: ''
		}
	},
	methods: {
		handleSelect(item){
			this.$emit('input', item)
		}
	},
	computed: {
		tableData(){
			return this.$store.getters['accounts/all'];
		}
	},
	created() {
		this.loading = true;
		this.$store.dispatch('accounts/fetch')
		.then(() => {
			this.fetchError = ''
		})
		.finally(()=>{
			this.loading = false;
		})
		.catch(err => {
			console.error(err.toJSON())
			this.fetchError = 'Can\'t reach the server'
			if(err.response && err.response.data.message){
				this.fetchError += `: [ERROR]: ${err.response.data.message}`
			} else {
				this.fetchError += `: [ERROR]: ${err.toJSON().message}`
			}
		})
	},
}
</script>

<style>

</style>