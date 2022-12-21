import { Button, Modal } from 'antd';
import React, { useState } from 'react';

const App: React.FC = () => {
    const [isModalOpen, setIsModalOpen] = useState(false);

    return (
        <div className="bg-background min-w-screen min-h-screen">
            <div className="flex items-center">
                <Button onClick={() => setIsModalOpen(true)}>Le button noir</Button>
                <Button onClick={() => setIsModalOpen(true)}>Le button blanc</Button>
                <Modal
                    open={isModalOpen}
                    onOk={() => setIsModalOpen(false)}
                    onCancel={() => setIsModalOpen(false)}
                    okButtonProps={{className: 'bg-primary hover:bg-tertiary'}}
                    title="A Modal"
                    okText="Okie dokie"
                >
                    <Button>Why is there another button ???</Button>
                </Modal>
            </div>
        </div>
    );
};

export default App;
