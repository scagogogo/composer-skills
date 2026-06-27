import React from 'react'
import { useTranslation } from 'react-i18next'
import { Tabs, Typography } from 'antd'
import SectionTitle from '../components/SectionTitle'
import CodeBlock from '../components/CodeBlock'

const { Title } = Typography

const QuickStartSection: React.FC = () => {
  const { t } = useTranslation()

  const tabItems = [
    {
      key: 'packagist',
      label: t('quickStart.tabPackagist'),
      children: <CodeBlock code={t('quickStart.packagistCode')} />,
    },
    {
      key: 'composer',
      label: t('quickStart.tabComposer'),
      children: <CodeBlock code={t('quickStart.composerCode')} />,
    },
    {
      key: 'autoInstall',
      label: t('quickStart.tabAutoInstall'),
      children: <CodeBlock code={t('quickStart.autoInstallCode')} />,
    },
  ]

  return (
    <section id="quickStart">
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('quickStart.title')} subtitle={t('quickStart.subtitle')} />

        <div style={{ maxWidth: 800, margin: '0 auto 32px' }}>
          <Title level={5}>{t('quickStart.install')}</Title>
          <CodeBlock code={t('quickStart.installCode')} language="bash" />
        </div>

        <div style={{ maxWidth: 900, margin: '0 auto' }}>
          <Tabs
            defaultActiveKey="packagist"
            items={tabItems}
            centered
            style={{ marginBottom: 40 }}
          />
        </div>

        <div style={{ maxWidth: 800, margin: '0 auto' }}>
          <Title level={4}>{t('quickStart.convenienceTitle')}</Title>
          <CodeBlock code={t('quickStart.convenienceCode')} />
        </div>
      </div>
    </section>
  )
}

export default QuickStartSection
