import {ref} from "vue";


export default function useLoading() {
    const loading = ref(false)
    const setLoading = (val: boolean) => {
        loading.value = val
    }
    return {loading, setLoading}
}