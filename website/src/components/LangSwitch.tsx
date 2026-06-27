import React from 'react'
import { useTranslation } from 'react-i18next'
import { Button } from 'antd'
import { GlobalOutlined } from '@ant-design/icons'

const LangSwitch: React.FC = () => {
  const { i18n } = useTranslation()
  const currentLang = i18n.language

  const toggleLang = () => {
    const newLang = currentLang === 'en' ? 'zh' : 'en'
    i18n.changeLanguage(newLang)
    localStorage.setItem('lang', newLang)
  }

  return (
    <Button
      size="small"
      icon={<GlobalOutlined />}
      onClick={toggleLang}
      style={{
        fontWeight: 600,
        borderRadius: 8,
        border: '1px solid #E2E8F0',
        color: '#475569',
        fontSize: 13,
      }}
    >
      {currentLang === 'en' ? '中文' : 'EN'}
    </Button>
  )
}

export default LangSwitch
