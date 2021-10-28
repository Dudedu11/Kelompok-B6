import React from 'react'
import { StyleSheet, Text, View,ScrollView } from 'react-native'
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
        
        <ScrollView>
            <View style={[styles.content]}>
            <View style={[styles.box]}></View>
            <View style={[styles.box]}></View>
            <View style={[styles.box]}></View>
            <View style={[styles.box]}></View>
            <View style={[styles.box]}></View>
            <View style={[styles.box]}></View>
            <View style={[styles.box]}></View>
            <View style={[styles.box]}></View>
            <View style={[styles.box]}></View>
            <View style={[styles.box]}></View>
            <View style={[styles.box]}></View>
            <View style={[styles.box]}></View>
          </View>
        </ScrollView>
        <View style={[styles.header]}>
          <Text style={styles.sectionTitle}>Feed</Text>

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
      backgroundColor: '#FFFFFF',
      borderBottomColor : '#A9A9A9',
      zIndex: 10
    },
    content: {
      alignItems: 'center',
      marginTop: 50,
      marginBottom: 40,
      backgroundColor: '#F8F8FF'
    },
    footer: {
      height: 70,
      position: 'absolute',
      left: 0,
      right: 0,
      bottom: 0,
      borderRadius: 50,
      backgroundColor: '#1E90FF'
  
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
    sectionTitleBottom : {
      fontSize: 24,
      fontWeight: '600',
      marginBottom: 12,
      marginLeft: 30,
      color: Colors.black
  
    }
  });