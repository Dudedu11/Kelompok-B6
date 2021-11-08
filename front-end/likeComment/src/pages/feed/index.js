import React from 'react'
import { StyleSheet, Text, Button, Alert, SafeAreaView, View,ScrollView, Image, TouchableNativeFeedback } from 'react-native'
import {
 Message,
 Post,
 LikeFalse,
 LikeTrue
} from '../../assets'

import {
    Header,
    LearnMoreLinks,
    Colors,
    DebugInstructions,
    ReloadInstructions,
  } from 'react-native/Libraries/NewAppScreen';
const Feed = () => {
    return (
        <View style={styles.container}>
        Like : nama_variable[0]['like'];
        <ScrollView>
            <View style={[styles.content]}>
            <View style={[styles.box]}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={[styles.box]}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={[styles.box]}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={[styles.box]}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={[styles.box]}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={[styles.box]}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={[styles.box]}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={[styles.box]}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={[styles.box]}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={[styles.box]}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={[styles.box]}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={[styles.box]}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
          </View>
        </ScrollView>
        <View style={[styles.header]}>
        <Text style={styles.sectionTitle}>Feed
        </Text>
        <Message height = {24}  marginTop = {15} marginLeft = {295} position = 'absolute' />  
        </View>
       </View>
    )
}

export default Feed

const styles = StyleSheet.create({
    container: {
      flex: 1,
      flexDirection: 'column',
      justifyContent: 'center'
    },
    header: {
      height: 50,
      position: 'absolute',
      left: 0,
      right: 0,
      top: 0,
      bottom : 10,
      backgroundColor: '#FFFFFF',
      borderBottomColor : 'F8F8FF',
    },
    content: {
      alignItems: 'center',
      marginTop: 50,
      marginBottom: 40,
      backgroundColor: '#F8F8FF'
    },
    box: {
      width: 330,
      height: 340,
      backgroundColor: '#FFFFFF',
      borderRadius: 30,
      marginBottom: 10,
    },
    sectionTitle: {
      fontSize: 24,
      fontWeight: '600',
      marginTop: 12,
      marginLeft: 30,
      color: Colors.black
    },
  });