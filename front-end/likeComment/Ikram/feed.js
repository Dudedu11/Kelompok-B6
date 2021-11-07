import React from 'react'
import { StyleSheet, Text, View, Button, ScrollView, Image } from 'react-native'
import Post from '../assets/icons/post.png';
import Message from '../assets/icons/message.svg';
import {
    Header,
    LearnMoreLinks,
    Colors,
    DebugInstructions,
    ReloadInstructions,
  } from 'react-native/Libraries/NewAppScreen';
import Like from '../components/Like';

const feed = () => {
    return (
        <View style={styles.container}>
        
        {/* <ScrollView>
            <View style={styles.content}>
            <View style={styles.box}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={styles.box}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={styles.box}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={styles.box}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={styles.box}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={styles.box}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={styles.box}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={styles.box}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={styles.box}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={styles.box}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={styles.box}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
            <View style={styles.box}>
            <Post height = {210}  marginTop = {10} marginLeft = {5} position = 'absolute'  /> 
            </View>
          </View>
        </ScrollView> */}
        {/* <ScrollView>
            <View>
            <Text>asuu</Text>
            </View>
        </ScrollView> */}
        <View style={styles.header}>
        <Text style={styles.sectionTitle}>Feed</Text>
        <Image 
            style={{marginTop:15, marginLeft:350, position:'absolute'}}
            source={require('../assets/icons/message.png')} />
        {/* <Message /> */}
        </View>
        <ScrollView>
            <View style={styles.content}>
                <View style={styles.box}>
                    <Image 
                    style={{width: 306.85, height:231.8, marginTop:20, marginLeft:10, position:'absolute'}}
                    source={require('../assets/icons/post.png')} />
                    <Like/>
                </View>
                <View style={styles.box}>
                    <Image 
                    style={{width: 306.85, height:231.8, marginTop:20, marginLeft:10, position:'absolute'}}
                    source={require('../assets/icons/post.png')} />
                    <Like/>
                </View>
                <View style={styles.box}>
                    <Image 
                    style={{width: 306.85, height:231.8, marginTop:20, marginLeft:10, position:'absolute'}}
                    source={require('../assets/icons/post.png')} />
                    <Like/>
                </View>
                <View style={styles.box}>
                    <Image 
                    style={{width: 306.85, height:231.8, marginTop:20, marginLeft:10, position:'absolute'}}
                    source={require('../assets/icons/post.png')} />
                    <Like/>
                </View>
            </View>
        </ScrollView>
       </View>
    )
}

export default feed

// const styles = StyleSheet.create({
//     container: {
//         flex: 1,
//         alignItems: 'center',
//         justifyContent: 'center',
//         backgroundColor: '#8fccbbc'
//     }
// })

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
      bottom : 0,
      backgroundColor: '#FFFFFF',
      borderBottomColor : 'F8F8FF',
      zIndex: 6,
    },
    content: {
      alignItems: 'center',
      marginTop: 50,
      marginBottom: 40,
      backgroundColor: '#F8F8FF',
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
