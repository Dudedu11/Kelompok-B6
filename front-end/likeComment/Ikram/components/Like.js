import React, { Component } from 'react';
// import type {Node} from 'react';
import {
  SafeAreaView,
  ScrollView,
  StatusBar,
  StyleSheet,
  Text,
  TouchableOpacity,
  useColorScheme,
  View,
} from 'react-native';

export default class Like extends Component {

  state={
    toggle:true
  }

  _onPress(){
    const newState = !this.state.toggle;
    this.setState({toggle:newState})
  }

  render(){
    const {toggle} = this.state;
    const textValue = toggle?"🤍":"❤️";
    const textValue2 = toggle?"   1,545":"   1,546";
    return (
      <View style={styles.container}>
  
        <View>
          <TouchableOpacity 
          onPress={()=>this._onPress()}
          style={{margin:20, marginTop:250, flex:1, height: 60, backgroundColor: '', }}>
            <Text style={{marginTop: 20, fontSize:16, }}>
                {textValue}
                <Text style={{color:'black', fontSize:12, }}>
                    {textValue2}
                </Text>
            </Text>
            
          </TouchableOpacity>
        </View>
        
        
      
      </View>
    );  
  }
  
};


const styles = StyleSheet.create({
    container: {
        flex: 1,
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'flex-start',
        // backgroundColor: 'whitesmoke',
    },
});