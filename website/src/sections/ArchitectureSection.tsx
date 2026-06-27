import React from 'react'
import { useTranslation } from 'react-i18next'
import { Table, Image } from 'antd'
import SectionTitle from '../components/SectionTitle'

const ArchitectureSection: React.FC = () => {
  const { t } = useTranslation()

  const layers = t('architecture.layers', { returnObjects: true }) as Array<{
    layer: string
    func: string
    pkg: string
  }>

  const columns = [
    {
      title: t('architecture.layerHeader'),
      dataIndex: 'layer',
      key: 'layer',
      width: '25%',
      render: (text: string) => <strong>{text}</strong>,
    },
    {
      title: t('architecture.functionHeader'),
      dataIndex: 'func',
      key: 'func',
      width: '45%',
    },
    {
      title: t('architecture.packageHeader'),
      dataIndex: 'pkg',
      key: 'pkg',
      width: '30%',
      render: (text: string) => (
        <code style={{ background: '#F1F5F9', padding: '1px 6px', borderRadius: 2, fontSize: 13 }}>
          {text}
        </code>
      ),
    },
  ]

  return (
    <section id="architecture" className="section-alt">
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('architecture.title')} subtitle={t('architecture.subtitle')} />

        <div style={{ textAlign: 'center', marginBottom: 32 }}>
          <Image
            src={`${import.meta.env.BASE_URL}images/architecture.png`}
            alt="Architecture"
            style={{ maxWidth: '90%' }}
            preview={false}
          />
        </div>

        <Table
          columns={columns}
          dataSource={layers.map((item, i) => ({ ...item, key: i }))}
          pagination={false}
          bordered
        />
      </div>
    </section>
  )
}

export default ArchitectureSection
