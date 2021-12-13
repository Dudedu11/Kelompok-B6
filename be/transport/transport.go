package transport

import (
	"be/datastruct"
	"be/logging"
	"be/service"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Melike(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Acce ss-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var suka datastruct.Menyukai
	var member datastruct.Member

	err := json.NewDecoder(r.Body).Decode(&suka)
	if err != nil {
		log.Fatalf("Tidak bisa mendecode dari request body.  %v", err)
	}
	
	getApiMember := fmt.Sprintf("https://61a9ef35bfb110001773efaa.mockapi.io/api/Member/%d", suka.User_id)
	response, _ := http.Get(getApiMember)
	if response.StatusCode != 200 {
		w.WriteHeader(response.StatusCode) 
		w.Write([]byte("Not found"))
		return
	} 

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &member)

	getApiFeed := fmt.Sprintf("https://61a9ef35bfb110001773efaa.mockapi.io/api/Feed/%d", suka.Feed_id)
	response_, _ := http.Get(getApiFeed)
	if response.StatusCode != 200 {
		w.WriteHeader(response_.StatusCode) 
		w.Write([]byte("Not found"))
		return
	} 


	suka.Tgl_like = logging.GetDateTimeNowInString()

	service.MLike(suka)

	logging.Log(fmt.Sprintf("user dengan id %d menyukai feed dengan id %d", suka.User_id, suka.Feed_id))

	res := datastruct.Response1{
		ID_user: suka.User_id,
		Name_user: member.Username,
		ID_feed: suka.Feed_id,
		Is_like: suka.Status_like,
		Message:  "Berhasil memberikan like",
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(res)
}

func TmplknSemuaLike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["feed_id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	likes, err := service.ShowAllLike(int64(id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data buku. %v", err)
	}

	var response datastruct.Response2
	response.Status = 200
	response.Message = "Success"
	response.Data = likes

	json.NewEncoder(w).Encode(response)
}

func TmplknLike(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var getAPIMember string
	var post datastruct.Feed

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["feed_id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	jmlLike  := service.ShowLike(int64(id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	
	getAPIMember = fmt.Sprintf("https://61a9ef35bfb110001773efaa.mockapi.io/api//Feed/%d", id)

	response, _ := http.Get(getAPIMember)
		if response.StatusCode != 200 {
			w.WriteHeader(response.StatusCode)
			w.Write([]byte("Not found"))
			return
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(responseData, &post)

		logging.Log(fmt.Sprintf("%d menampilkan data like", id))

	response3:= datastruct.Response3{
		Status: 200,
		Message: "Success",
		JumlahLike: jmlLike,
	}
	
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response3)
}

func UptLike(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	params := mux.Vars(r)

	id_user_, err := strconv.Atoi(params["user_id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	id_feed_, err := strconv.Atoi(params["feed_id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	var suka datastruct.Menyukai

	err = json.NewDecoder(r.Body).Decode(&suka)

	if err != nil {
		log.Fatalf("Tidak bisa decode request body.  %v", err)
	}

	updatedRows := service.UpdateLike(int64(id_user_), int64(id_feed_), suka)

	msg := fmt.Sprintf("Data telah berhasil diupdate. Status yang diupdate %v rows/record", updatedRows)

	res := datastruct.Response4{
		Status: 200,
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

////////////////////////////BACKEND KOMENNNNN ///////////////////////////////////

func TmbhKomen(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Acce ss-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var komen datastruct.Komen
	var member datastruct.Member

	err := json.NewDecoder(r.Body).Decode(&komen)
	if err != nil {
		log.Fatalf("Tidak bisa mendecode dari request body.  %v", err)
	}
	
	getApiMember := fmt.Sprintf("https://61a9ef35bfb110001773efaa.mockapi.io/api/Member/%d", komen.User_id)
	response, _ := http.Get(getApiMember)
	if response.StatusCode != 200 {
		w.WriteHeader(response.StatusCode) 
		w.Write([]byte("Not found"))
		return
	} 

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &member)

	getApiFeed := fmt.Sprintf("https://61a9ef35bfb110001773efaa.mockapi.io/api/Feed/%d", komen.Feed_id)
	response_, _ := http.Get(getApiFeed)
	if response.StatusCode != 200 {
		w.WriteHeader(response_.StatusCode) 
		w.Write([]byte("Not found"))
		return
	} 


	komen.Post_comment = logging.GetDateTimeNowInString()

	service.TambahKomen(komen)

	logging.Log(fmt.Sprintf("user dengan id %d mengomentari feed dengan id %d", komen.User_id, komen.Feed_id))

	res := datastruct.Response5{
		ID_user: komen.User_id,
		Name_user: member.Username,
		ID_feed: komen.Feed_id,
		Message:  "Berhasil tambah komen",
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(res)
}

func TmplknSemuaKomen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["feed_id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	likes, err := service.TampilAllKomen(int64(id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data komen. %v", err)
	}

	var response datastruct.Response6
	response.Status = 200
	response.Message = "Success"
	response.Data = likes

	json.NewEncoder(w).Encode(response)
}
