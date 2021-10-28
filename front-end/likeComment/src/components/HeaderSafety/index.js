import React from 'react'
import { StyleSheet, Text, View, TouchableOpacity } from 'react-native'
import { IconBell, IconManage, IconPesan } from '../../assets'
import { useNavigation } from '@react-navigation/native';

const HeaderSafety = () => {
    const navigation = useNavigation();
    return (
        <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center' }}>
            <Text>Toolbar</Text>
        </View>
    )
}

export default HeaderSafety

const styles = StyleSheet.create({})